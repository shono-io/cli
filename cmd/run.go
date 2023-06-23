package cmd

import (
	"fmt"
	"github.com/shono-io/shono/local"
	"github.com/shono-io/shono/runtime"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an artifact",
	Long:  `Retrieve and run the given artifact`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg runtime.RunConfig
		c, err := cmd.Flags().GetString("config")
		if err != nil {
			fmt.Println(err)
		}
		if c != "" {
			b, err := os.ReadFile(c)
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to read config file %q: %v", c, err))
			}
			if err := yaml.Unmarshal(b, &cfg); err != nil {
				fmt.Println(fmt.Sprintf("failed to unmarshal config file %q: %v", c, err))
			}
		}

		u, err := cmd.Flags().GetString("artifact")
		if err != nil {
			fmt.Println(err)
		}

		artifact, err := local.LoadArtifact(u)
		if err != nil {
			logrus.Errorf("failed to load artifact: %v", err)
			return
		}

		if err := runtime.RunArtifact(cfg, artifact); err != nil {
			logrus.Errorf("failed to run artifact: %v", err)
		}
	},
}

func init() {
	artifactCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("artifact", "a", "", "the artifact to run")
	runCmd.Flags().StringP("config", "c", "", "the path to the runtime configuration file")
}
