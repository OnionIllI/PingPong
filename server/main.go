package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func wsHandler(w http.ResponseWriter, r *http.Request) {

}
