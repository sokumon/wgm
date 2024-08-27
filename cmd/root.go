package cmd

import (
	"fmt"
	"os"
	"wgm/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wgm",
	Short: "CLI tool to manage Wireguard Server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			utils.ExecuteCommand("echo", "Hello World")
		}
	},
}



func Execute() {
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
