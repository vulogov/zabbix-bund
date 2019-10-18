package cmd

import (
  "github.com/spf13/cobra"
  log "github.com/sirupsen/logrus"
  bctx "github.com/vulogov/zabbix-bund/bund_context"
)

// helloCmd represents the hello command
var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "cache server",
	Long:  `Caches metrics in memory DB`,
	Run: func(cmd *cobra.Command, args []string) {
    if bctx.InstanceN == "" {
      bctx.InstanceN = "cache"
    }
    log.Debug("Executing ", bctx.InstanceN, "#", bctx.InstanceNo)
  },
}

func init() {
	rootCmd.AddCommand(cacheCmd)
}
