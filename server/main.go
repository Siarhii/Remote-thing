package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"

	"server/handlers"
)

func main() {

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow the specific origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"}, // Allowed methods
		AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Allowed headers
		AllowCredentials: true, // Allow credentials like cookies or headers
	})
	
	http.HandleFunc("/connect", handlers.ConnectRouteHandler)
	http.HandleFunc("/api/getdeviceID",handlers.GetDeviceIDHandler)
	http.HandleFunc("/api/sendcommand",handlers.SendCommandHandlerr)
	http.HandleFunc("/api/devices",handlers.FetchDevicesHandler)

	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", corsHandler.Handler(http.DefaultServeMux))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}


