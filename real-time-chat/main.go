package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request){
	conn,_:=upgrader.Upgrade(w,r,nil)

	client:=&Client{
		hub:hub,
		conn:conn,
		send: make(chan []byte, 256),
	}

	client.hub.register<-client

	go client.writePump()
	go client.readPump()
}

func main() {
    hub := NewHub()
    go hub.Run()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(hub, w, r)
    })

    // ADD THIS LINE
    fmt.Println("Server starting on :8080...")

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}