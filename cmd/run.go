package cmd

import (
	"fmt"
	"github.com/shono-io/shono/local"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an artifact",
	Long:  `Retrieve and run the given artifact`,
	Run: func(cmd *cobra.Command, args []string) {
		u, err := cmd.Flags().GetString("artifact")
		if err != nil {
			fmt.Println(err)
		}

		artifact, err := local.LoadArtifact(u)
		if err != nil {
			logrus.Errorf("failed to load artifact: %v", err)
		}

		if err := artifact.Run(); err != nil {
			logrus.Errorf("failed to run artifact: %v", err)
		}
	},
}

func init() {
	artifactCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("artifact", "a", "", "the artifact to run")
}
