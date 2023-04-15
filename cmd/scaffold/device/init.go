/*
Copyright © 2023 Jonathan Taylor <jonathan.taylor@cuanschutz.edu>
*/

package device

import (
	"github.com/spf13/cobra"
	"scaffold/node"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new device",
	Long:  `Generate a new device folder structure`,
	Run: func(cmd *cobra.Command, args []string) {
		n, err := cmd.PersistentFlags().GetString("name")
		if err != nil {
			panic(err)
		}
		d, err := cmd.PersistentFlags().GetString("dest")
		if err != nil {
			panic(err)
		}
		node.NewDevice(n, d)
	},
}

func init() {
	deviceCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().String("name", "newDevice", "device name")
	initCmd.PersistentFlags().StringP("dest", "d", ".", "specify the device location (defaults to current directory)")
}
