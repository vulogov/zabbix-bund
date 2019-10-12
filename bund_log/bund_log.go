package bund_log
import (
  "os"
  "github.com/op/go-logging"
)

var Log = logging.MustGetLogger("bund")
var format = logging.MustStringFormatter(
        `%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func Init_Log() {
  log_backend := logging.NewLogBackend(os.Stderr, "", 0)
  log_fmt :=  logging.NewBackendFormatter(log_backend, format)
  log_level := logging.AddModuleLevel(log_backend)
  log_level.SetLevel(logging.ERROR, "")
  logging.SetBackend(log_fmt)
}
