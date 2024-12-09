package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	globalvariables "server/globalVariables" //my file name and packname is exact same so its importing wth an alias? import numpy as np like thing
	"server/helpers"
	"server/types"
	"time"
)

type CommandRequest struct {
	DeviceID string `json:"deviceId"`
	Command string `json:"Command"`
	Timer string `json:"scheduleTime"`
	Password string `json:"password"`
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

	if request.DeviceID == "" || request.Command == "" {
		fmt.Printf("Device id : %v and device command : %v",request.DeviceID,request.Command)
		http.Error(w, "Missing deviceId or command", http.StatusBadRequest)
		return
	}

	DeviceID := request.DeviceID
	if !helpers.CheckIfDeviceRegistered(DeviceID){
		http.Error(w, fmt.Sprintf("Device with ID %s is not registered", DeviceID), http.StatusNotFound)
		return
	}

	wsn, exists := globalvariables.LiveWebSocketConnectionsMap[DeviceID]
	if !exists {
		http.Error(w, fmt.Sprintf("Device with ID %s is Offline", DeviceID), http.StatusNotFound) //offline as in not have active socket connection with server
		return
	}

	if request.Command != "Shutdown" && request.Command != "Sleep" && request.Command != "Restart"{
		http.Error(w,  fmt.Sprintf("Invalid Command : %v", request.Command), http.StatusBadRequest)
		return
	}

	if !helpers.VerifyTimer(request.Timer){
		http.Error(w,  fmt.Sprintf("Invalid timer : %v", request.Timer), http.StatusBadRequest)
		return
	}

	if !helpers.VerifyDevicePassword(DeviceID, request.Password) {
		fmt.Printf("Here3 : %v\n",request.Timer)
		fmt.Printf("Here4 : %v\n",request.Password)
		http.Error(w, "Device Password is not correct", http.StatusUnauthorized)
		return
	}

	msg := types.Message{
		Event:   "Command",
		Content: request.Command + "_" + request.Timer, 
	}

	wsn.WriteChan <- msg

	select {
	case resMsg := <-wsn.CommandResponseChan:
		fmt.Printf("Response from client for Command is: %v\n", resMsg)
	case <-time.After(4 * time.Second):
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
