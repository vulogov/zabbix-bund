package cmd

import (
	"github.com/spf13/cobra"
  log "github.com/sirupsen/logrus"
  lisp "github.com/vulogov/zabbix-bund/bund_lisp"
  bctx "github.com/vulogov/zabbix-bund/bund_context"
)

var (

)

// helloCmd represents the hello command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "execute shell",
	Long:  `Execute bund internal shell`,
	Run: func(cmd *cobra.Command, args []string) {
      lisp.ZB_lisp_init()
      switch bctx.ScriptFile {
         case "_":
           log.Debug("Executing REPL")
           lisp.ZB_repl()
         case "-":
           log.Debug("Executing code from STDIN")
         default:
           log.Debug("Executing code from ",bctx.ScriptFile)
      }
	},
}

func init() {
  shellCmd.PersistentFlags().StringVarP(&bctx.ScriptFile, "file", "f", "_", "Execute script with embedded LISP or provide access to REPL (default is REPL)")
  shellCmd.PersistentFlags().StringVarP(&bctx.LispConfig, "cfg", "", "", "Configuration for the embedded LISP")
  rootCmd.AddCommand(shellCmd)
}
