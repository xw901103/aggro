package main

import (
  "net/http"
  "github.com/gorilla/websocket"
  "log"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
  conn *websocket.Conn
  buff chan []byte
}

func (c *Client) pump() {
  defer func() {
    c.conn.Close()
  }()
  for {
    _, message, err := c.conn.ReadMessage()
    if err != nil {
      log.Println(err)
      break
    }
    log.Println(message)
  }
}

func serveHoleWebSocket(w http.ResponseWriter, r *http.Request) {
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Println(err)
    return
  }
  client := &Client {conn: conn, buff: make(chan []byte, 1024)}
  go client.pump()
}
