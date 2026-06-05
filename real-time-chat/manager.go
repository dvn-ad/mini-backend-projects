package main

import (
	"encoding/json"
	"fmt"
	"time"
) 


type Hub struct {
    // The Map: Keys are Client pointers, Values are booleans
    clients map[*Client]bool

    // Inbound messages from the clients
    broadcast chan []byte

    // Register requests from the clients
    register chan *Client

    // Unregister requests from clients
    unregister chan *Client
}

func NewHub() *Hub {
    return &Hub{
        broadcast:  make(chan []byte),
        register:   make(chan *Client),
        unregister: make(chan *Client),
        clients:    make(map[*Client]bool),
    }
}

func (h *Hub) Run(){
	for {
		select{
		case client := <-h.register:
			h.clients[client]=true
			msg := Message{
				Type:      "system",
				Sender:    "System",
				Content:   fmt.Sprintf("%s joined the chat", client.username),
				Timestamp: time.Now().Format(time.RFC3339),
			}
			if marshalled, err := json.Marshal(msg); err == nil {
				for c := range h.clients {
					select {
					case c.send <- marshalled:
					default:
						close(c.send)
						delete(h.clients, c)
					}
				}
			}
		case client:=<-h.unregister:
			if _,ok:=h.clients[client];ok{
				delete(h.clients,client)
				close(client.send)
				
				msg := Message{
					Type:      "system",
					Sender:    "System",
					Content:   fmt.Sprintf("%s left the chat", client.username),
					Timestamp: time.Now().Format(time.RFC3339),
				}
				if marshalled, err := json.Marshal(msg); err == nil {
					for c := range h.clients {
						select {
						case c.send <- marshalled:
						default:
							close(c.send)
							delete(h.clients, c)
						}
					}
				}
			}
		case message:=<-h.broadcast:
			for client:=range h.clients{
				if h.clients[client]==true{
					client.send<-message
				}
			}
		}
	}
}