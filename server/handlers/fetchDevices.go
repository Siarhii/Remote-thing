package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	globalvariables "server/globalVariables"
	"server/helpers"
)

func FetchDevicesHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var devices []map[string]interface{}

	helpers.UpdateDeviceOnlineStatus() //i am updating online or offline status of device during every fetch device req which i think might be inefficient but i dont know a better approach to show live devices to user in kindof realtime 
	for connectionCode, device := range globalvariables.AddedDeviceMap {

		if !device.ClientAdded {
			continue
		}

		status := "offline"
		if device.Online {
			status = "online"
		}

		onlineSince := fmt.Sprintf("%d mins ago", device.OnlineSince)

		var scheduledAction interface{}
		if device.ScheduledAction {
			scheduledAction = map[string]interface{}{
				"type":          device.Command, 
				"remainingTime": device.Timer,   
			}
		}

		deviceInfo := map[string]interface{}{
			"id":              connectionCode,
			"name":            device.DeviceName,
			"status":          status,
			"onlineSince":     onlineSince,
			"scheduledAction": scheduledAction,
		}

		// Append the device information to the devices slice
		devices = append(devices, deviceInfo)
	}

	// Set the response header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the devices slice as JSON and send it to the frontend
	if err := json.NewEncoder(w).Encode(devices); err != nil {
		fmt.Printf("JSON Encoding Error: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
