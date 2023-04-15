/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package device

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"scaffold/modbus"
	"scaffold/node/goClient"
	"scaffold/node/yaml"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client is used to generate a device API client in the chosen language",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		device := yaml.NewYAMLDevice("127.0.0.1", 8081, modbus.DefaultClient(logger))
		device.Load("./nodes")
		srv := goClient.NewService("https://127.0.0.1:8081")
		for _, n := range device.Nodes {
			out, err := os.Create(fmt.Sprintf("./clients/%s_client.go", n.Node))
			if err != nil {
				zap.Error(err)
			}
			err = srv.Flush(out, n)
			if err != nil {
				zap.Error(err)
			}
		}
	},
}

func init() {
	nodeCmd.AddCommand(clientCmd)
	clientCmd.PersistentFlags().StringP("lang", "-l", "go", "the language to build the client for")
}
