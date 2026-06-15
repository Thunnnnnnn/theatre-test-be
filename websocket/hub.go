package websocket

import "github.com/gorilla/websocket"

type Hub struct {
	Clients map[*websocket.Conn]bool
}

var WSHub = &Hub{
	Clients: make(map[*websocket.Conn]bool),
}
