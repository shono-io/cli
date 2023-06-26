package cmd

import (
	"github.com/spf13/cobra"
)

// artifactCmd represents the artifact command
var artifactCmd = &cobra.Command{
	Use:   "artifact",
	Short: "Perform action against artifacts",
	Long: `Artifacts are executable units of code that can be run by shono.
The contain logic and optionally tests.

3 types of artifacts are supported:
- concept artifacts: artifacts that react to events focussed on a specific concept
- injector artifacts: artifacts that read from a source and inject events into the backbone
- extractor artifacts: artifacts that read from the backbone and write to a target
`,
}

func init() {
	rootCmd.AddCommand(artifactCmd)
}
