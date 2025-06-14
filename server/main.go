package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for simplicity; adjust as needed
	},
}

var connections = make(map[string]*websocket.Conn)

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func assignRole(conn *websocket.Conn) string {
	switch len(connections) {
	case 0:
		return "left"
	case 1:
		return "right"
	default:
		return ""
	}
}

func sendMessageToPlayer(role string, message string, conn *websocket.Conn) {
	conn.WriteJSON(map[string]string{
		role: message,
	})
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	log.Println("New WebSocket connection established")

	role := assignRole(conn)

	if role == "" {
		log.Println("Room full, rejecting connection")
		conn.WriteJSON(map[string]string{
			"type":   "error",
			"reason": "room full",
		})

		conn.Close()
		return
	}

	connections[role] = conn

	log.Printf("New connection: role=%s, ip=%s\n", role, r.RemoteAddr)

	sendMessageToPlayer(role, "You are connected as "+role, conn)

}
