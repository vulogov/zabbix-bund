package bund_lisp

import (
  log "github.com/sirupsen/logrus"
  "github.com/glycerine/zygomys/zygo"
  bctx "github.com/vulogov/zabbix-bund/bund_context"

)

var (
  lcfg = zygo.NewZlispConfig("bund")
)

func ZB_lisp_init() {
  log.Debug("Embedded LISP initialized")
  log.Debug("Version : ", zygo.Version())
  lcfg.DefineFlags()
  err := lcfg.ValidateConfig()
  if err != nil {
    log.Panic(err)
  }
  lcfg.Prompt = bctx.PS1
}
