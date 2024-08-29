package cmd


import (
	"fmt"
	"github.com/spf13/cobra"
	"wgm/utils"
)


var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Creates a client.conf for the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Installing Wireguard")
		if utils.CheckOS() == "Ubuntu" || utils.CheckOS() == "Debian" {
			utils.ExecuteCommand("sudo","apt","update")
			utils.ExecuteCommand("sudo","apt","install","wireguard","-y")
			fmt.Print("Installed Wireguard successfully")
		}
	},
}