package sockets

import (
	"fmt"
	"time"

	"server/types"
)

func StartPingMessages(wsc *types.WebSocketConnection){
	pingMessage := types.Message{
		Event : "Ping",
		Content: "i can add date/time here",
	}
	fmt.Print("Sending ping")
	for {
		wsc.WriteChan <- pingMessage

		time.Sleep(1 * time.Second)
	}
}