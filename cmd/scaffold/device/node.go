/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package device

import (
	"fmt"
	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "node is used to create new nodes or get them from the internet",
	Long:  `I haven't implemented getting them from the internet yet, and this function by itself does nothing :)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented")
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)
}
