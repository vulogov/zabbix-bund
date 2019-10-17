package bund_lisp

import (
  log "github.com/sirupsen/logrus"
  "github.com/glycerine/zygomys/zygo"
)
func ZB_repl() {
  log.Debug("Entering embedded LISP REPL")
  zygo.ReplMain(lcfg)
  //Zenv.LoadString("(+ 40 2)")
  //expr, _ := Zenv.Run()
  //fmt.Println(expr)

}
