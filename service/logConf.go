package service

import (
  "encoding/xml"
  "io/ioutil"
  "log"
  _"fmt"
)

type Result struct {
  Service Service `xml:Service`
}

type Service struct {
  Status string 
  Level string 
  Token string
}

func LogConf(client string)(serviceConf Service){
  content, err := ioutil.ReadFile("config/client/" + client + ".xml")
  if err != nil {
    log.Fatal(err)
  }
  var result Result
  err = xml.Unmarshal(content, &result)
  if err != nil {
    log.Fatal(err)
  }
  return result.Service
}