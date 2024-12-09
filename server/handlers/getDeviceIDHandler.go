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

	var deviceDetail types.DeviceDetail
	if err := json.NewDecoder(r.Body).Decode(&deviceDetail); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if deviceDetail.DeviceName == "" || deviceDetail.DevicePassword == "" {
		http.Error(w, "DeviceName and Password are required", http.StatusBadRequest)
		return
	}

	// later i will generate the connection code based on username and device name + randomcode
	// connectionCode := helpers.GenerateRandomCode(5) + "_" + deviceDetail.DeviceName
	connectionCode := helpers.GenerateRandomCode(5) 

	deviceExists := false
	for _, device := range globalvariables.AddedDeviceMap {
		if device.DeviceName == deviceDetail.DeviceName {
			deviceExists = true
			break
		}
	}
	if deviceExists {
		http.Error(w, "Device with same name exists.", http.StatusBadRequest)
		return
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
