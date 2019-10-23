package cmd

import (
  "os"
	"fmt"
  "github.com/vulogov/zabbix-bund/bund_log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

  bctx "github.com/vulogov/zabbix-bund/bund_context"
  log "github.com/sirupsen/logrus"
  bund "github.com/vulogov/zabbix-bund/bund_logic"
)

var (
	// Used for flags.
	rootCmd = &cobra.Command{
		Use:   "zabbix-bund",
		Short: "Universal client/server for zabbix-bund",
		Long: `zb - is a universal application for building distributed metrics collection
and processing system.`,
	}
)

// Execute executes the root command.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&bctx.CfgFile, "config", "", "config file (default is $HOME/.zabbix-bund)")
	rootCmd.PersistentFlags().StringVarP(&bctx.Logverbose, "verbose", "v", "info", "Level for the logging (trace,debug,warning,info,fatal)")
  rootCmd.PersistentFlags().StringVarP(&bctx.Logoutput, "logfmt", "l", "text", "Format of the log output (text,json)")
  rootCmd.PersistentFlags().StringVarP(&bctx.InstanceN, "name", "n", "", "Name of the instance")
  rootCmd.PersistentFlags().Uint32VarP(&bctx.InstanceNo, "number", "#", 0, "Number of the instance")
  rootCmd.PersistentFlags().StringVarP(&bctx.DataDir, "data", "d", "", "Data directory")
  rootCmd.PersistentFlags().StringVarP(&bctx.RaftDir, "rdata", "a", "", "Raft directory")
  rootCmd.PersistentFlags().StringVarP(&bctx.HTTPBind, "http", "t", ":21080", "HTTP endpoint bind address")
  rootCmd.PersistentFlags().StringVarP(&bctx.RaftBind, "raft", "r", ":21081", "RAFT endpoint bind address")
  rootCmd.PersistentFlags().StringVarP(&bctx.JoinAddr, "join", "j", "localhost:21081", "RAFT join address")
  rootCmd.PersistentFlags().Bool("is_raft", true, "Do start the RAFT")
  rootCmd.PersistentFlags().Bool("is_rest", true, "Do start the REST")

}

func initConfig() {
	if bctx.CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(bctx.CfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Error(err.Error())
      os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".zabbix-bund")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Debug(fmt.Sprintf("Using config file:", viper.ConfigFileUsed()))
	}
  bund_log.Init_Log(bctx.Logverbose, bctx.Logoutput)
  is_raft, _ := rootCmd.PersistentFlags().GetBool("is_raft")
  is_rest, _ := rootCmd.PersistentFlags().GetBool("is_rest")
  bund.Init_Internal_Components(is_raft, is_rest)
  log.Debug("root_init process complete")
}
