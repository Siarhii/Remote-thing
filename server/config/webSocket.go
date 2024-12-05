package config

import (
	"net/http"

	"github.com/gorilla/websocket"
)


var Upgrader = websocket.Upgrader{
	ReadBufferSize:  8192,
    WriteBufferSize: 8192, //will have to change later
	CheckOrigin: func(r *http.Request) bool {
		return true 
	},
}