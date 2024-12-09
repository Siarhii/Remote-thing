package helpers

import (
	"fmt"
	globalvariables "server/globalVariables"
	"time"
)

func CleanupDeviceMap() {
	for {
		// Sleep for 10 minutes before running cleanup again
		time.Sleep(10 * time.Minute)

		// Create a list of keys to delete after the iteration to avoid modifying the map during iteration
		var toDelete []string

		// Iterate over the map and check the ClientAdded flag
		for code, device := range globalvariables.AddedDeviceMap {
			if !device.ClientAdded {
				toDelete = append(toDelete, code)
			}
		}
		
		for _, code := range toDelete {
			fmt.Printf("Deleting device with code %s, name %s\n", code, globalvariables.AddedDeviceMap[code].DeviceName)
			delete(globalvariables.AddedDeviceMap, code)
		}
	}
}