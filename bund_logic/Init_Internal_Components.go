package bund_logic

import (
  log "github.com/sirupsen/logrus"
  brft "github.com/vulogov/zabbix-bund/bund_raft"
)

func Init_Internal_Components(is_raft bool, is_rest bool) {
  log.Debug("Initializing internal components", is_raft, is_rest)
  if is_raft == true {
    brft.Raft_Init()
  }
}
