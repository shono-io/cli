package cmd

import (
	"github.com/shono-io/cli/generator"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build [path]",
	Short: "Generate artifacts based on the given path",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := generator.Generate(args[0]); err != nil {
			logrus.Errorf("failed to build artifact: %v", err)
		}
	},
}

func init() {
	ArtifactCmd.AddCommand(buildCmd)
}
