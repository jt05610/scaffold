/*
Copyright © 2023 Jonathan Taylor <jonrtaylor12@gmail.com>
*/

package device

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"scaffold/modbus"
	"scaffold/node/yaml"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs the device server",
	Long:  `The device should have nodes configured`,
	Run: func(cmd *cobra.Command, args []string) {
		host, err := cmd.PersistentFlags().GetString("address")
		if err != nil {
			panic(err)
		}
		port, err := cmd.PersistentFlags().GetInt("port")
		if err != nil {
			panic(err)
		}
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		device := yaml.NewYAMLDevice(host, port, modbus.DefaultClient(logger))
		nodes, err := cmd.PersistentFlags().GetString("nodes")
		if err != nil {
			panic(err)
		}
		device.Load(nodes)
		device.Serve()
	},
}

func init() {
	deviceCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().StringP("nodes", "n", "", "node config directory")
	serveCmd.PersistentFlags().IntP("port", "p", 8081, "node config directory")
	serveCmd.PersistentFlags().StringP("address", "a", "127.0.0.1", "node config directory")

}
