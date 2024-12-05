package types

import (
	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	Conn  *websocket.Conn
	DeviceID string
	UserID string
	DeviceName string
	WriteChan chan Message
	CommandResponseChan chan Message
	StatsResponseChan chan Message
	ErrChan chan error
	DoneChan chan struct {}
	OnlineInMinutes uint32 
}


func NewWebSocketConnection(conn *websocket.Conn,deviceID string,userID string) *WebSocketConnection {
	return &WebSocketConnection{
		Conn : conn,
		DeviceID: deviceID,
		UserID: userID,
		DeviceName : "Not Avail for now",
		CommandResponseChan: make(chan Message,1),
		StatsResponseChan: make(chan Message,2),
		WriteChan: make(chan Message,5),
		ErrChan: make(chan error,2),
		DoneChan: make(chan struct{}),
		OnlineInMinutes: 0,
	}
}