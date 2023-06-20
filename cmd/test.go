package cmd

import (
	"github.com/shono-io/shono/local"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long:  `Retrieve and test the given artifact`,
	Run: func(cmd *cobra.Command, args []string) {
		u, err := cmd.Flags().GetString("artifact")
		if err != nil {
			logrus.Error(err)
		}

		artifact, err := local.LoadArtifact(u)
		if err != nil {
			logrus.Errorf("failed to load artifact: %v", err)
		}

		if err := artifact.Test(logrus.TraceLevel.String()); err != nil {
			logrus.Errorf("failed to test artifact: %v", err)
		}
	},
}

func init() {
	artifactCmd.AddCommand(testCmd)
	testCmd.Flags().StringP("artifact", "a", "", "the artifact to run")
}
