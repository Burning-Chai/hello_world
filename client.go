package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	send   chan *message
}

func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			c.socket.WriteJSON(msg)
		} else {
			break
		}
	}
	c.socket.Close()
}
