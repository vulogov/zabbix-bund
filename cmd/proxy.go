package cmd

import (
  "github.com/spf13/cobra"
  log "github.com/sirupsen/logrus"
  bctx "github.com/vulogov/zabbix-bund/bund_context"
)

// helloCmd represents the hello command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "proxy server",
	Long:  `Runs the metrics collectors`,
	Run: func(cmd *cobra.Command, args []string) {
    if bctx.InstanceN == "" {
      bctx.InstanceN = "proxy"
    }
    log.Debug("Executing ", bctx.InstanceN, "#", bctx.InstanceNo)
  },
}

func init() {
	rootCmd.AddCommand(proxyCmd)
}
