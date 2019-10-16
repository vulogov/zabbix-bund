package cmd

import (
  "os"
	"fmt"
  "github.com/vulogov/zabbix-bund/bund_log"
  log "github.com/sirupsen/logrus"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
  logverbose  string
  logoutput   string


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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zabbix-bund)")
	rootCmd.PersistentFlags().StringVarP(&logverbose, "verbose", "v", "info", "Level for the logging (trace,debug,warning,info,fatal)")
  rootCmd.PersistentFlags().StringVarP(&logoutput, "logfmt", "l", "text", "Format of the log output (text,json)")

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
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
  bund_log.Init_Log(logverbose, logoutput)
  log.Debug("root_init process complete")
}
