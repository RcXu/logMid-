package model

import(
  _"fmt"
  "service"
  "log"
)

type Object struct {
  Service string
  Token string
  Loglevel string
  Data []string
}

func LogToFile(client string, token string, Loglevel string, Data []string){
  conf := service.LogConf(client)

  //身份校验失败
  if  token != conf.Token {
    return 
  }

  //关闭日志
  if conf.Status == "0" {
    return
  }

  //判断日志级别，0=>其他日志 1=> warning日志 2=>error 日志
  if Loglevel < conf.Level {
    return
  }
  
  log.Println(Data)
}
