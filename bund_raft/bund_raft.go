package bund_raft

import (
  "os"
  "path"
  "fmt"
  "time"
  "net"
  log "github.com/sirupsen/logrus"
  bctx "github.com/vulogov/zabbix-bund/bund_context"
  "github.com/hashicorp/raft"
  raftboltdb "github.com/hashicorp/raft-boltdb"
)

var (
  Cfg = raft.DefaultConfig()
  FullPath  string
  RaftDb    raftboltdb.BoltStore
  SnapStore raft.SnapshotStore
  Transport raft.NetworkTransport
)

func Raft_Init() {
	log.Debug("Initialize RAFT")
  if bctx.JoinAddr == "" {
    log.Info("This will be the first node in the cluster")
    Cfg.EnableSingleNode = true
  } else {
    log.Info("Attempts to join existing cluster will be made")
    Cfg.EnableSingleNode = false
  }
  if bctx.DataDir == "" {
    log.Fatal("Empty --data parameter is not permitted")
  }
  log.Debug("Base directory is: ", bctx.DataDir)
  log.Debug("Instance name is: ", bctx.InstanceN)
  if bctx.InstanceN == "" {
    FullPath = path.Join(bctx.DataDir, bctx.BundId.URN())
  } else {
    FullPath = path.Join(bctx.DataDir, bctx.InstanceN, fmt.Sprint(bctx.InstanceNo))
  }
  log.Debug("Data direcory is ", FullPath)
  err := os.MkdirAll(FullPath, 0700)
  if err != nil {
    log.Fatal(err)
  }
  RaftDb, err := raftboltdb.NewBoltStore(path.Join(FullPath, "raft_db"))
  if err != nil {
		log.Fatal(err)
	}
	SnapStore, err := raft.NewFileSnapshotStore(FullPath, 1, os.Stdout)
  if err != nil {
		log.Fatal(err)
	}

  log.Debug("RAFT binding: ", bctx.RaftBind)
  log.Debug("RAFT advertisement: ", bctx.AdvertiseAddr)
  localAddr, err := net.ResolveTCPAddr("tcp", bctx.AdvertiseAddr)
  if err != nil {
    log.Fatal(err)
  }
  Transport, err := raft.NewTCPTransport(bctx.RaftBind, localAddr, 3, 5*time.Second, os.Stdout)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(RaftDb, SnapStore, Transport)
}
