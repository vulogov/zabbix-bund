package cmd

import (
  "github.com/spf13/cobra"
  log "github.com/sirupsen/logrus"
  bctx "github.com/vulogov/zabbix-bund/bund_context"
)

var arbiterCmd = &cobra.Command{
	Use:   "arbiter",
	Short: "raft arbiter",
	Long:  `First process in the cluster`,
	Run: func(cmd *cobra.Command, args []string) {
    if bctx.InstanceN == "" {
      bctx.InstanceN = "arbiter"
    }
    log.Debug("Executing ", bctx.InstanceN, "#", bctx.InstanceNo)
  },
}

func init() {
	rootCmd.AddCommand(arbiterCmd)
}
