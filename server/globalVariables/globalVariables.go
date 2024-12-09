package globalvariables

import (
	"server/types"
)

var LiveWebSocketConnectionsMap = make(map[string] *types.WebSocketConnection)
//contains
// key will be deviceID and value is
// type WebSocketConnection struct {
// 	Conn  *websocket.Conn
// 	DeviceID string
// 	ReadChan chan Message
// 	WriteChan chan Message
// 	ErrChan chan error
// 	DoneChan chan struct {}
// }

var UserToDeviceIDMap = make(map[string]string)
//for now key:value is  userName:connectionCode he generated //maybe i dont need it now

var AddedDeviceMap = make(map[string] *types.DeviceDetail)
//key:value is connectionCode:deviceDetails