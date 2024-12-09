package sockets

import (
	"time"

	"server/types"
)

func StartPingMessages(wsc *types.WebSocketConnection){
	pingMessage := types.Message{
		Event : "Ping",
		Content: "i can add date/time here",
	}
	for {
		wsc.WriteChan <- pingMessage

		time.Sleep(1 * time.Second)
	}
}