package cmd

import (
  "os"
	"fmt"
  "github.com/vulogov/zabbix-bund/bund_log"
  bctx "github.com/vulogov/zabbix-bund/bund_context"
  log "github.com/sirupsen/logrus"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

		// Search config in home directory with name ".zabbix-bund" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".zabbix-bund")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
  bund_log.Init_Log(bctx.Logverbose, bctx.Logoutput)
  log.Debug("root_init process complete")
}
