package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	send   chan *message
	uuid string
}

func (c *client) read() {
	for {
		if messageType, m, err := c.socket.ReadMessage(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(messageType)
			//fmt.Println(string(m))

			var msg *message
			if err := json.Unmarshal(m, &msg); err != nil {
				//fmt.Println(err)
			} else {
				c.uuid = msg.Sender
				c.socket.WriteJSON(msg)
			}
		}
	}
	c.socket.Close()
}
