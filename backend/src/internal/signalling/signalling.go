package signalling

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/VivekHalder/webRTC/handlers"
	"github.com/VivekHalder/webRTC/models"
	"github.com/gorilla/websocket"
)

// websocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: true,
}

var (
	rooms      = make(map[string][]*websocket.Conn)
	roomsMutex sync.Mutex
)

func addClientToRoom(roomID string, client *websocket.Conn) {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()
	rooms[roomID] = append(rooms[roomID], client)
}

func removeClientFromRoom(roomID string, client *websocket.Conn) {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	for i, c := range rooms[roomID] {
		if c == client {
			rooms[roomID] = append(rooms[roomID][:i], rooms[roomID][i+1:]...)
			break
		}
	}
}

func broadcastMessage(roomID string, msg models.Message, sender *websocket.Conn) {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	for _, c := range rooms[roomID] {
		if c != sender {
			if err := c.WriteJSON(msg); err != nil {
				log.Println("Write Error")
			}
		}
	}
}

func UpgraderToWebSockets(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		http.Error(w, "Error occured while upgrading the connection", http.StatusBadRequest)
		return
	}

	fmt.Println("Connection upgraded to websockets")
	handlers.AddClient(conn)
	defer func() {
		conn.Close()
		handlers.RemoveClient(conn)
	}()

	roomID := r.URL.Query().Get("roomID")
	if roomID == "" {
		msg := models.Message{Type: "error", Payload: "Meeting ID is required to join the meet."}
		conn.WriteJSON(msg)
		return
	}

	addClientToRoom(roomID, conn)
	defer removeClientFromRoom(roomID, conn)

	// handle the websocket and the message
	for {
		// how to message (struct)
		var msg models.Message
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println("Error reading the message: ", err)
			break
		}

		switch msg.Type {
		case "offer", "answer", "ice_candidates":
			broadcastMessage(roomID, msg, conn)

		default:
			log.Println("Unknown message type: ", msg.Type)
			errMsg := models.Message{Type: "error", Payload: "Unknown message type"}
			conn.WriteJSON(errMsg)
		}
	}
}
