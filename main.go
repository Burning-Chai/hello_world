package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

func main() {
	http.HandleFunc("/room", ServeHTTP)

	log.Println("Webサーバーを開始します。ポート:5000")
	if err := http.ListenAndServe("5000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func ServeHTTP(w http.ResponseWriter, req *http.Request) {

	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	client := &client{
		socket: socket,
		send:   make(chan *message, messageBufferSize),
	}

	client.read()
}
