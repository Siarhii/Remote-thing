package sockets

import (
	"encoding/json"
	"fmt"

	"server/helpers"
	"server/types"

	"github.com/gorilla/websocket"
)

func WriteToSocket(wsc *types.WebSocketConnection){
	for {	
		select {
			case <- wsc.DoneChan :
				return
			default :
				msg := <-wsc.WriteChan //this will block the code until a value is pushed into the writeChan so the loop wont be looping unless a value is pushed into writeChan
				switch msg.Event {
					case "Ping" , "Command" , "Stat" :
						//ignore as we dont need to do any addtional thing in writing message part
						//we are creating types.Message somewhere else and sending them to channel
					default:
						wsc.ErrChan <- fmt.Errorf("unexpected event type: %s in (WriteToSocket)", msg.Event)
						continue
				}

				msgBytes, err := json.Marshal(msg) 
				if helpers.CheckError(err,"during Marshal of write message (writeToSocket)") {
					wsc.ErrChan <- fmt.Errorf("during Marshal of write message event : %s (writeToSocket) ", msg.Event)
					//skip during marshal error
					continue
				}

				err = wsc.Conn.WriteMessage(websocket.TextMessage, msgBytes)
				if helpers.CheckError(err,"during Conn.WriteMessage (writeToSocket)") {
					wsc.ErrChan <- fmt.Errorf("during Conn.WriteMessage event : %s (writeToSocket)", msg.Event)
					CloseConnection(wsc)
					return
				}
				if(msg.Event == "Ping"){
					wsc.OnlineInMinutes += 1
				}
		}
	}
}