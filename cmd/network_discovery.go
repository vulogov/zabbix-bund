package cmd

import (
  "github.com/spf13/cobra"
  log "github.com/sirupsen/logrus"
  bctx "github.com/vulogov/zabbix-bund/bund_context"
)

// helloCmd represents the hello command
var ndiscCmd = &cobra.Command{
	Use:   "network-discovery",
	Short: "network-discovery server",
	Long:  `Discover nodes on the network`,
	Run: func(cmd *cobra.Command, args []string) {
    if bctx.InstanceN == "" {
      bctx.InstanceN = "network-discovery"
    }
    log.Debug("Executing ", bctx.InstanceN, "#", bctx.InstanceNo)
  },
}

func init() {
	rootCmd.AddCommand(ndiscCmd)
}
