/*
Copyright © 2023 Jonathan Taylor <jonrtaylor12@gmail.com>
*/

package device

import (
	"github.com/spf13/cobra"
	"os"
	"path"
	"scaffold/device"
	"scaffold/node/yaml"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create the project structure of a new node",
	Long:  `Generate a project structure where you will define and store all code related to a new node`,
	Run: func(cmd *cobra.Command, args []string) {
		srv := yaml.NewYAMLService()
		n, err := cmd.PersistentFlags().GetString("name")
		if err != nil {
			panic(err)
		}
		k, err := cmd.PersistentFlags().GetString("kind")
		if err != nil {
			panic(err)
		}
		home, err := os.UserHomeDir()
		p := path.Join(home, "scaffold", "nodes")
		err = os.MkdirAll(p, 0777)
		if err != nil {
			panic(err)
		}
		device.NewNode(n, k, p, srv)
	},
}

func init() {
	nodeCmd.AddCommand(newCmd)
	newCmd.PersistentFlags().String("kind", "software", "The kind of node. Can be software or hardware")
	newCmd.PersistentFlags().String("name", "newNode", "The name of your node. Example: linear_axis ")
}
