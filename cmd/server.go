package cmd

import (
  "github.com/spf13/cobra"
  log "github.com/sirupsen/logrus"
  bctx "github.com/vulogov/zabbix-bund/bund_context"
)

// helloCmd represents the hello command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "bund server",
	Long:  `Runs the server which computes incoming metrics against set of rules`,
	Run: func(cmd *cobra.Command, args []string) {
    if bctx.InstanceN == "" {
      bctx.InstanceN = "server"
    }
    log.Debug("Executing ", bctx.InstanceN, "#", bctx.InstanceNo)
  },
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
