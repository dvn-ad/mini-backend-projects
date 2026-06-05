package main

import (
	"github.com/gorilla/websocket"
	"encoding/json"
	"time"
)

type Message struct{
	Type		string `json:"type"`
	Sender		string `json:"sender"`
	Content		string `json:"content"`
	Timestamp	string `json:"timestamp"`
}

type Client struct {
    hub  *Hub
    conn *websocket.Conn // This comes from the gorilla/websocket library
    send chan []byte     // The personal mailbox
	username string
}

func (c *Client) readPump () {
	defer func(){
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for{
		_,message, err:= c.conn.ReadMessage()
		if err!=nil{
			break
		}
		msg := Message{
			Type: "chat",
			Sender: c.username,
			Content: string(message),
			Timestamp: time.Now().Format(time.RFC3339),
		}
		marshalled,_:=json.Marshal(msg)
		c.hub.broadcast<-marshalled
	}
}

func (c *Client) writePump () {
	defer func(){
		c.conn.Close()
	}()
	for{
		message,ok:=<-c.send
		if !ok{
			c.conn.WriteMessage(websocket.CloseMessage,[]byte{})
			return
		}
		c.conn.WriteMessage(websocket.TextMessage, message)
	}
}