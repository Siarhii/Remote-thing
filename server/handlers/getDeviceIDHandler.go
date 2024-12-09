package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	globalvariables "server/globalVariables"
	"server/helpers"
	"server/types"
)

func GetDeviceIDHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // struct to capture only the device name and password from the request
    var request struct {
        DeviceName     string `json:"DeviceName"`
        DevicePassword string `json:"DevicePassword"`
    }

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if request.DeviceName == "" || request.DevicePassword == "" {
        http.Error(w, "DeviceName and Password are required", http.StatusBadRequest)
        return
    }

    connectionCode := helpers.GenerateRandomCode(5)

    // Check if a device with the same name already exists
    for _, device := range globalvariables.AddedDeviceMap {
        if device.DeviceName == request.DeviceName {
            http.Error(w, "Device with same name exists.", http.StatusBadRequest)
            return
        }
    }

    deviceDetail := types.DeviceDetail{
        UserID:          "anon",              
        DeviceName:      request.DeviceName,    
        DevicePassword:  request.DevicePassword,
        ClientAdded:     false,                
        Online:          false,                 
        ScheduledAction: false,                 
        Command:         "none",                
        Timer:           "none",                
        OnlineSince:     0,                     
    }
    
    globalvariables.AddedDeviceMap[connectionCode] = &deviceDetail
   
    response := map[string]string{
        "connectionCode": connectionCode,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(response); err != nil {
        fmt.Printf("JSON Encoding Error: %v\n", err)
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}
