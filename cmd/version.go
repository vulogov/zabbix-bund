package cmd

import (
	"fmt"
  "github.com/fatih/color"
  "github.com/vulogov/zabbix-bund/bund_context"
	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "bund version",
	Long:  `Displays current version of the zabbix-bund package`,
	Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(color.YellowString(bund_context.Logo.String()))
		fmt.Println(color.WhiteString("Zabbix-Bund"), " ", color.RedString(bund_context.Version_Num), " ",color.GreenString(bund_context.Version_Release))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
