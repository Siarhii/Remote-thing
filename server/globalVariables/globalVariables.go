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
// 	UserID string
// 	ReadChan chan Message
// 	WriteChan chan Message
// 	ErrChan chan error
// 	DoneChan chan struct {}
// }

var UserToDeviceIDMap = make(map[string]string)
//for now key:value is  userName:DeviceID he generated //maybe i dont need it now