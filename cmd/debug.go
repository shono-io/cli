package cmd

import (
	"github.com/shono-io/shono/local"
	"github.com/shono-io/shono/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var debugCmd = &cobra.Command{
	Use:   "debug [artifact]",
	Short: "Debug an artifact",
	Long: `Start a shono-enhanced benthos instance to start processing data.
The input and output defined as part of the artifact will be replaced with stdin and stdout respectively.
An in-memory storage will be used to store the data
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		debug, _ := cmd.Flags().GetBool("debug")

		logLevel := "INFO"
		if debug {
			logLevel = "DEBUG"
		}

		artifact, err := local.LoadArtifact(args[0])
		if err != nil {
			logrus.Errorf("failed to load artifact: %v", err)
			return
		}

		if err := runtime.DebugArtifact(artifact, logLevel); err != nil {
			logrus.Errorf("failed to run artifact: %v", err)
		}
	},
}

func init() {
	ArtifactCmd.AddCommand(debugCmd)
	debugCmd.Flags().BoolP("debug", "d", false, "print verbose debug information")
}
