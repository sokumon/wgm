package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Gets you the version of the cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version 1")
	},
}
