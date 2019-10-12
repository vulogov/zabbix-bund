package cmd

import (
  "os"
	"fmt"
  "github.com/vulogov/zabbix-bund/bund_log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	Verbose     bool


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
		bund_log.Log.Error(err.Error())
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zabbix-bund)")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	//viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

	//rootCmd.AddCommand(addCmd)
	//rootCmd.AddCommand(initCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			bund_log.Log.Error(err.Error())
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
}
