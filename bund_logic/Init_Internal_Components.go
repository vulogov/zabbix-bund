package bund_logic

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  brft "github.com/vulogov/zabbix-bund/bund_raft"
  "github.com/google/uuid"
  bctx "github.com/vulogov/zabbix-bund/bund_context"
)

func Init_Internal_Components(is_raft bool, is_rest bool) {
  log.Debug("Initializing internal components")
  bctx.BundId, _  = uuid.NewUUID()
  if bctx.InstanceNo == 0 {
    bctx.InstanceNo = bctx.BundId.ID()
  }
  log.Debug(fmt.Sprintf("UUID: %s", bctx.BundId.URN()))
  if is_raft == true {
    brft.Raft_Init()
  }
}
