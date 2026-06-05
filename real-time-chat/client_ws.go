package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
    hub  *Hub
    conn *websocket.Conn // This comes from the gorilla/websocket library
    send chan []byte     // The personal mailbox
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
		c.hub.broadcast<-message
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