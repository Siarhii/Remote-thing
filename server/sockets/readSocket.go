package sockets

import (
	"encoding/json"
	"fmt"

	"server/helpers"
	"server/types"
)

func ReadFromSocket(wsc *types.WebSocketConnection){
	for {
		select {
			case <- wsc.DoneChan :
				return //close the loop if DoneChan is closed (indicator to tell us wether socket connection is live or not) 
			default:
				_,messageData,err := wsc.Conn.ReadMessage()
				if(helpers.CheckError(err,"during reading socket message (readFromSocket)")){
					wsc.ErrChan <- fmt.Errorf("during reading socket message (readFromSocket)")
					CloseConnection(wsc)
					return
				}

				var messageRecieved types.Message
				err = json.Unmarshal(messageData , &messageRecieved)
				if helpers.CheckError(err,"during unmarshal of read message (readFromSocket) ") {
					wsc.ErrChan <- fmt.Errorf("during unmarshal of read message (readFromSocket)")
					CloseConnection(wsc)
					return
				}

				switch messageRecieved.Event {
					case "Pong":
						//ignore
					case "CommandResponse" :
						wsc.CommandResponseChan <- messageRecieved
					case "StatResponse" :
						wsc.StatsResponseChan <- messageRecieved
					default:
						wsc.ErrChan <- fmt.Errorf("unexpected event type: %s in (readFromSocket)", messageRecieved.Event)
				}
		}
	}
}