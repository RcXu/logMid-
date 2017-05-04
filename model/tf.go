package model

import (
  "fmt"
  "net"
  "net/rpc"
  "scribe"
  "github.com/samuel/go-thrift/thrift"
)

type scribeServiceImplementation int

func (s *scribeServiceImplementation) Log(messages []*scribe.LogEntry) (scribe.ResultCode, error) {
  for _, m := range messages {
    LogToFile(m.Client, m.Token, m.Loglevel, m.Message)
  }
  return scribe.ResultCodeOk, nil
}

func TFservice() {
  scribeService := new(scribeServiceImplementation)
  rpc.RegisterName("Thrift", &scribe.ScribeServer{Implementation: scribeService})
  ln, err := net.Listen("tcp", ":8181")
  if err != nil {
    panic(err)
  }
  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Printf("ERROR: %+v\n", err)
      continue
    }
    fmt.Printf("New connection %+v\n", conn)
    t := thrift.NewTransport(thrift.NewFramedReadWriteCloser(conn, 0), thrift.BinaryProtocol)
    go rpc.ServeCodec(thrift.NewServerCodec(t))
  }
}
