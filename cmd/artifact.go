package cmd

import (
	"github.com/spf13/cobra"
)

// artifactCmd represents the artifact command
var artifactCmd = &cobra.Command{
	Use:   "artifact",
	Short: "Artifact Operations",
	Long:  `Artifacts can be interpreted and run or tested using the cli.`,
}

func init() {
	rootCmd.AddCommand(artifactCmd)
}
