package handlers

import (
	"fmt"
	"net/http"

	"server/config"
	globalvariables "server/globalVariables"
	"server/helpers"
	"server/sockets"
	"server/types"
)

func ConnectRouteHandler(w http.ResponseWriter, r *http.Request) {
fmt.Print("AJHHHA")
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method. Only GET is allowed.", http.StatusMethodNotAllowed)
		return
	}

	queryParams := r.URL.Query()
	deviceID := queryParams.Get("deviceID")
	userID := queryParams.Get("userID")

	if deviceID == "" || userID == "" || len(userID) > 5 || len(deviceID) > 5 {
		http.Error(w, "Missing or invalid query parameters: deviceID and userID.", http.StatusBadRequest)
		return
	}

	var upgrader = config.Upgrader
	conn, err := upgrader.Upgrade(w, r, nil)

	if helpers.CheckError(err,"upgrading to websocket") {
		http.Error(w, "Failed to establish WebSocket connection", http.StatusInternalServerError)
	}

	fmt.Println("Client connected!") ////////////////////////////////////////////Temp

	websocketConnection := types.NewWebSocketConnection(conn,deviceID,userID)
	defer sockets.CloseConnection(websocketConnection) // will close connection and also delete key:value from LiveWebSocket... map
	globalvariables.LiveWebSocketConnectionsMap[deviceID] = websocketConnection
	
	go sockets.ReadFromSocket(websocketConnection)

	go sockets.WriteToSocket(websocketConnection)

	go sockets.StartPingMessages(websocketConnection)

	
	select {}
}









