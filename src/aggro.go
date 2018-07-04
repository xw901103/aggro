package main

import (
  "fmt"
  "encoding/json"
  "io/ioutil"
  "os"
  "net/http"
  //"github.com/lib/pq"
  "log"
)

type RuntimeConfig struct {
  Hostname string `json:"hostname"`
  Port int `json:"port"`
}

func getRuntimeConfig(path string) (config RuntimeConfig, err error) {
  raw, err := ioutil.ReadFile(path)
  if err == nil {
    json.Unmarshal(raw, &config)
  }
  return
}

func main() {
  config, err := getRuntimeConfig("./config.json")
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  http.HandleFunc("/", viewHandler)
  http.HandleFunc("/hole", func(w http.ResponseWriter, r *http.Request) {
    serveHoleWebSocket(w, r)
  })
  log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Hostname, config.Port), nil))
}
