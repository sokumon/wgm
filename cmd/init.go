package cmd


import (
	"fmt"
	"github.com/spf13/cobra"
)


var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a wg.conf for the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initilizing WireGuard")
	},
}