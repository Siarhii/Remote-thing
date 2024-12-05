package main

import (
	"fmt"
	"net/http"

	"server/handlers"
)

// var upgrader = config.Upgrader

// var devicesStatus = make(map[string]string) // {"deviceCode" : "online/offline"}
// var userDevices = make(map[string][]string) // {"user" : [deviceCode1,deviceCode2]}
// var deviceSocketConnectionMap = make(map[string]*websocket.Conn)

// func sendCommandHandler(w http.ResponseWriter, r *http.Request){
// 	go sockets.sendCommand(deviceCode string,userID string,command string,timer uint64){}
// }

func main() {

	http.HandleFunc("/connect", handlers.ConnectRouteHandler)
	// http.HandleFunc("/api/sendCommand",sendCommandHandler())
	http.HandleFunc("/api/allotdeviceID",handlers.GetDeviceIDHandler)

	http.HandleFunc("/api/sendcommand",handlers.SendCommandHandlerr)

	//starting server
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil) 
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

