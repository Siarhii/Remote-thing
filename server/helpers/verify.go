package helpers

import (
	"strconv"

	globalvariables "server/globalVariables"
)

func VerifyDevicePassword(deviceID string,password string) bool {
	devicePassword := globalvariables.AddedDeviceMap[deviceID].DevicePassword

	if devicePassword == password {
		return true
	} else {
		return false
	}
}

func VerifyTimer(timer string) bool {
	
	minutes, err := strconv.Atoi(timer)
	if err != nil {
		return false
	}

	if minutes > -1 && minutes < 44000 {
		return true
	}

	return false
}
