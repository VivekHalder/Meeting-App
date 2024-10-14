package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VivekHalder/webRTC/models"
)

func HandleOffer(w http.ResponseWriter, r *http.Request) {
	var msg models.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for client := range clients {
		if err := client.WriteJSON(msg); err != nil {
			client.Close()
			RemoveClient(client)
		}
	}
}
