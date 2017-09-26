package slacktest

import (
	"log"

	websocket "github.com/gorilla/websocket"
)

func sendToChannel(s string) {
	log.Printf("Got a message for the channel: %s", s)
	sendMessageChannel <- s
}

func handlePendingMessages(c *websocket.Conn) {
	for m := range sendMessageChannel {
		log.Printf("Got an incoming message: %s", m)
		err := c.WriteMessage(websocket.TextMessage, []byte(m))
		if err != nil {
			log.Printf("error writing message to websocket: %s", err.Error())
		}
		seenMessageChannel <- string(m)
	}
}
