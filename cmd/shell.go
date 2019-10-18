package cmd

import (
  "fmt"
	"github.com/spf13/cobra"
  log "github.com/sirupsen/logrus"
  lisp "github.com/vulogov/zabbix-bund/bund_lisp"
  bctx "github.com/vulogov/zabbix-bund/bund_context"
  "github.com/erikdubbelboer/gspt"
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
           if bctx.InstanceN == "" {
             bctx.InstanceN = "shell"
           }
           log.Debug("Executing REPL")
           gspt.SetProcTitle("[zabbix-bund] Interactive shell")
           lisp.ZB_repl()
         case "-":
           if bctx.InstanceN == "" {
             bctx.InstanceN = "lisp-stdin"
           }
           gspt.SetProcTitle(fmt.Sprintf("[zabbix-bund] LISP from stdin %s", bctx.InstanceN))
           log.Debug("Executing code from STDIN")
         default:
           if bctx.InstanceN == "" {
             bctx.InstanceN = bctx.ScriptFile
           }
           gspt.SetProcTitle(fmt.Sprintf("[zabbix-bund] LISP from file %s", bctx.InstanceN))
           log.Debug("Executing code from ",bctx.ScriptFile)
      }
	},
}

func init() {
  shellCmd.PersistentFlags().StringVarP(&bctx.ScriptFile, "file", "f", "_", "Execute script with embedded LISP or provide access to REPL (default is REPL)")
  rootCmd.AddCommand(shellCmd)
}
