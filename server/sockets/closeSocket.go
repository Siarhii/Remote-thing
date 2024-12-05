package sockets

import (
	globalvariables "server/globalVariables"
	"server/helpers"
	"server/types"
)

func CloseConnection(wsc *types.WebSocketConnection) {
    close(wsc.DoneChan) // This will notify other goroutinea about closing of connection,so if due to a write error the connection is closed,read loop will know and stop
    close(wsc.CommandResponseChan)
    close(wsc.StatsResponseChan)
    close(wsc.WriteChan)
	close(wsc.ErrChan)
	err := wsc.Conn.Close();
	helpers.CheckError(err,"Error closing WebSocket connection in (closeConnection)")


	delete(globalvariables.LiveWebSocketConnectionsMap , wsc.DeviceID)
}