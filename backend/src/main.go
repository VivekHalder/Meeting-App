package main

import (
	"log"
	"net/http"

	"github.com/VivekHalder/webRTC/internal/signalling"
)

func main() {
	http.HandleFunc("/ws", signalling.UpgraderToWebSockets)

	serverAddr := ":8080"
	log.Printf("Starting the server at the port %s.\n", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatalf("Failed to start the server: %v.", err)
	}
}
