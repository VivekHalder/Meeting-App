package handlers

import (
	"sync"

	"github.com/gorilla/websocket"
)

var (
	clients = make(map[*websocket.Conn]bool)
	mu      sync.Mutex
)

func AddClient(client *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	clients[client] = true
}

func RemoveClient(client *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	delete(clients, client)
}
