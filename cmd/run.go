package cmd

import (
	"fmt"
	"github.com/shono-io/shono/local"
	"github.com/shono-io/shono/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [artifact]",
	Short: "Run an artifact",
	Long: `Start a shono-enhanced benthos instance to start processing data.

Most likely, additional system information (like login credentials and connection urls) 
are required to run the artifact. This command will look for those in the following locations:
- ./systems.yaml
- ~/.shono/systems.yaml

Each location will be checked in order and the first one that is found will be used.
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appId, err := cmd.Flags().GetString("id")
		if err != nil {
			fmt.Println(err)
		}

		storageId, _ := cmd.Flags().GetString("storage")
		debug, _ := cmd.Flags().GetBool("debug")

		logLevel := "INFO"
		if debug {
			logLevel = "DEBUG"
		}

		systems, err := runtime.LoadSystems()
		if err != nil {
			fmt.Println(fmt.Sprintf("failed to load systems: %v", err))
		}

		artifact, err := local.LoadArtifact(args[0])
		if err != nil {
			logrus.Errorf("failed to load artifact: %v", err)
			return
		}

		cfg := runtime.RunConfig{
			ApplicationId:   appId,
			StorageSystemId: storageId,
		}

		if err := runtime.RunArtifact(cfg, systems, artifact, logLevel); err != nil {
			logrus.Errorf("failed to run artifact: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("id", "i", "", "the application id")
	runCmd.Flags().StringP("storage", "s", "", "the storage system id, only applicable if the artifact is a concept artifact")
	runCmd.Flags().BoolP("debug", "d", false, "print verbose debug information")
}
