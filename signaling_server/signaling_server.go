package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	peers      = make(map[string]*websocket.Conn) // Peer ID -> Connection
	peersMutex = sync.Mutex{}
	sfuConn    *websocket.Conn                     // SFU Connection
)

// HandleWebSocket for both clients and SFU
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// Read the initial message (peer ID or SFU flag)
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("Read error:", err)
		return
	}

	peerID := string(message)

	// If the connection is to the SFU, set the SFU connection
	if peerID == "sfu" {
		sfuConn = conn
		log.Println("SFU connected")
		return
	}

	// Handle peer connection
	peersMutex.Lock()
	peers[peerID] = conn
	peersMutex.Unlock()
	log.Printf("New peer connected: %s", peerID)

	// Remove peer on disconnect
	defer func() {
		peersMutex.Lock()
		delete(peers, peerID)
		peersMutex.Unlock()
		log.Printf("Peer %s disconnected", peerID)
	}()

	// Listen for incoming messages
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read error for %s: %v", peerID, err)
			break
		}

		log.Printf("Received message from %s: %s", peerID, msg)

		// If message starts with "sfu|", it's intended for the SFU
		if peerID != "sfu" {
			if sfuConn != nil {
				err := sfuConn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					log.Println("Failed to send to SFU:", err)
				}
			}
		} else {
			// If it's from the SFU, handle messages to peers
			for id, peerConn := range peers {
				if id != "sfu" {
					err := peerConn.WriteMessage(websocket.TextMessage, msg)
					if err != nil {
						log.Println("Failed to send to peer:", id)
					}
				}
			}
		}
	}
}

// Main function to start the server
func main() {
	http.HandleFunc("/ws", HandleWebSocket)

	addr := "localhost:8080"
	log.Printf("Signaling server started at %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
