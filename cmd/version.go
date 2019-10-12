package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "bund version",
	Long:  `Displays current version of the zabbix-bund package`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[BUND> 0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
