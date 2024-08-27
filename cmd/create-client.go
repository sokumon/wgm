package cmd


import (
	"fmt"
	"github.com/spf13/cobra"
)


var createCliCmd = &cobra.Command{
	Use:   "create-client",
	Short: "Creates a client.conf for the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create Client")
	},
}