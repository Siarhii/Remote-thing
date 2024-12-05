package handlers

import (
	"encoding/json"
	"net/http"
	"server/helpers"
)

func GetDeviceIDHandler (w http.ResponseWriter,_ *http.Request){

	deviceID := helpers.GenerateRandomCode(5)
	
	///bruhhhh save it in mongodb,the code generated will be based on username and a random code like deviceID_username 
	//which user will paste into his client app

	response := map[string]string{
		"deviceID": deviceID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) 
	

	if err := json.NewEncoder(w).Encode(response); 
	err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}