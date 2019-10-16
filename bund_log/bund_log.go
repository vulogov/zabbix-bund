package bund_log
import (
  log "github.com/sirupsen/logrus"
)


func Init_Log(verbose string, output string) {
  switch output {
    case "json":
      log.SetFormatter(&log.JSONFormatter{})
    default:
      log.SetFormatter(&log.TextFormatter{})
  }
  switch verbose {
    case "trace":
      log.SetLevel(log.TraceLevel)
    case "debug":
      log.SetLevel(log.DebugLevel)
    case "warning":
      log.SetLevel(log.WarnLevel)
    case "error":
      log.SetLevel(log.ErrorLevel)
    case "fatal":
      log.SetLevel(log.FatalLevel)
    default:
      log.SetLevel(log.InfoLevel)
  }
  log.WithFields(log.Fields{
    "output": output,
    "level": verbose,
  }).Info("Log output configured")
}
