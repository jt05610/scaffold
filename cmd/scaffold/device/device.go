/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package device

import (
	"fmt"
	"github.com/spf13/cobra"
)

// deviceCmd represents the device command
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "device is focused on managing devices",
	Long:  `use device to generate and manage scientific robotics projects`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("device called")
	},
}

func init() {
	rootCmd.AddCommand(deviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
