package bund_context

import (
  "github.com/common-nighthawk/go-figure"
  "github.com/google/uuid"
)

var Version = "Zabbix-Bund version 0.1.1 (development release)"
var Version_Num = "0.1.1"
var Version_Release = "(development release)"
var Logo = figure.NewFigure("[theBund> ", "o8", true)
var PS1 = "[BUND> "

var (
  CfgFile     string
  Logverbose  string
  Logoutput   string
  ScriptFile  string
  InstanceNo  uint32
  InstanceN   string
  DataDir     string
  RaftDir     string
  HTTPBind    string
  RaftBind    string
  JoinAddr    string
  BundId      uuid.UUID
)
