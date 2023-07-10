package cmd

import (
	"fmt"
	"github.com/shono-io/shono/local"
	"github.com/shono-io/shono/runtime"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"

	_ "github.com/benthosdev/benthos/v4/public/components/all"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test [artifact]",
	Short: "Run the tests for the given artifact",
	Long: `Every artifact has some logic associated with it. 

If together with the logic, also tests are provided, this command will run them.

Please keep in mind only tests you have written yourself will be run.`,
	Args: cobra.ExactArgs(1),
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

		artifact, err := local.LoadArtifact(args[0])
		if err != nil {
			logrus.Errorf("failed to load artifact: %v", err)
			return
		}

		if err := runtime.TestArtifact(artifact, "TRACE"); err != nil {
			logrus.Errorf("failed to run artifact: %v", err)
		}
	},
}

func init() {
	ArtifactCmd.AddCommand(testCmd)
	testCmd.Flags().StringP("artifact", "a", "", "the artifact to run")
	testCmd.Flags().StringP("config", "c", "", "the path to the runtime configuration file")
}
