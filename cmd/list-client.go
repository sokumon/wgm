package cmd


import (
	"fmt"
	"github.com/spf13/cobra"
)


var listcliCmd = &cobra.Command{
	Use:   "list-client",
	Short: "Creates a wg.conf for the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing wireguard")
	},
}