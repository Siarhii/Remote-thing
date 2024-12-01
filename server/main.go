package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true 
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Handle requests at the root URL ("/")
	http.HandleFunc("/", handler)

	// Start the server on port 8080
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
