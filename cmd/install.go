package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Creates a client.conf for the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing Wireguard")
	},
}
