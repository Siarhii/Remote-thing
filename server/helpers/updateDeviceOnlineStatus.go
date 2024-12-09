package helpers

import (
	globalvariables "server/globalVariables"
)

//to flip flop deviceDetails.Online as true or false depending on if there is a live socket connection between device and server
func UpdateDeviceOnlineStatus() {
	for deviceID, deviceDetail := range globalvariables.AddedDeviceMap {
		if _, exists := globalvariables.LiveWebSocketConnectionsMap[deviceID]; exists {
			deviceDetail.Online = true
		} else {
			deviceDetail.Online = false
		}
	}

	// // Print check
	// for deviceID, deviceDetail := range globalvariables.AddedDeviceMap {
	// 	fmt.Printf("DeviceID: %s, DeviceName: %s, Online: %v\n", deviceID, deviceDetail.DeviceName, deviceDetail.Online)
	// }
}
