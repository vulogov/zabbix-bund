package cmd

import (
	"github.com/spf13/cobra"
  log "github.com/sirupsen/logrus"
)

var (
  scriptFile string
)

// helloCmd represents the hello command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "execute shell",
	Long:  `Execute bund internal shell`,
	Run: func(cmd *cobra.Command, args []string) {
      switch scriptFile {
         case "_":
           log.Debug("Executing REPL")
         case "-":
           log.Debug("Executing code from STDIN")
         default:
           log.Debug("Executing code from ",scriptFile)
      }

	},
}

func init() {
  shellCmd.PersistentFlags().StringVarP(&scriptFile, "file", "f", "_", "Execute script with embedded LISP or provide access to REPL (default is REPL)")
	rootCmd.AddCommand(shellCmd)
}
