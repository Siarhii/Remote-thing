package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	globalvariables "server/globalVariables" //my file name and packname is exact same so its importing wth an alias? import numpy as np like thing
	"server/types"
	"time"
)

type CommandRequest struct {
	DeviceID string `json:"deviceId"`
}

func SendCommandHandlerr(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request CommandRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if request.DeviceID == "" {
		http.Error(w, "Missing deviceId", http.StatusBadRequest)
		return
	}

	DeviceID := request.DeviceID

	wsn, exists := globalvariables.LiveWebSocketConnectionsMap[DeviceID]
	if !exists {
		http.Error(w, fmt.Sprintf("Device with ID %s not found", DeviceID), http.StatusNotFound)
		return
	}

	msg := types.Message{
		Event:   "Command",
		Content: "Meowmeow", 
	}

	fmt.Printf("LETSSOOGOGO : %v \n",DeviceID)
	wsn.WriteChan <- msg
	fmt.Printf("fassssssssssssssst : %v\n",DeviceID)

	select {
	case resMsg := <-wsn.CommandResponseChan:
		fmt.Printf("Response from client for Command is: %v\n", resMsg)
	case <-time.After(10 * time.Second):
		http.Error(w, "Timeout waiting for response from device", http.StatusRequestTimeout)
		return
	}

	// Prepare the response to the client
	response := map[string]string{
		"status":   "success",
		"deviceId": request.DeviceID,
	}

	// Set the Content-Type header and send the response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
