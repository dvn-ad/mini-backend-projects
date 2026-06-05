package main

import(
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
		case client:=<-h.unregister:
			if _,ok:=h.clients[client];ok{
				delete(h.clients,client)
				close(client.send)
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