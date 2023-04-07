/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package device

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "scaffold",
	Version: "0.1.0",
	Short:   "Scaffold is used to create nodes and run devices",
	Long: `To initialize a robot project, run
	scaffold init -n <device_name>
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
