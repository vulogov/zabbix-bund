package main

import (
	"github.com/vulogov/zabbix-bund/cmd"
  "github.com/vulogov/zabbix-bund/bund_log"
)

func main() {
  bund_log.Init_Log()
	cmd.Execute()
}
