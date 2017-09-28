package slacktest

import (
	"context"
	"fmt"
	"log"
	"time"

	websocket "github.com/gorilla/websocket"
	slack "github.com/nlopes/slack"
)

func queueForWebsocket(s string) {
	seenOutboundMessages.Lock()
	seenOutboundMessages.messages = append(seenOutboundMessages.messages, s)
	seenOutboundMessages.Unlock()
	sendMessageChannel <- s

}

func handlePendingMessages(c *websocket.Conn) {
	for m := range sendMessageChannel {
		err := c.WriteMessage(websocket.TextMessage, []byte(m))
		if err != nil {
			log.Printf("error writing message to websocket: %s", err.Error())
			continue
		}
	}
}

func postProcessMessage(m string) {
	seenInboundMessages.Lock()
	seenInboundMessages.messages = append(seenInboundMessages.messages, m)
	seenInboundMessages.Unlock()
	// send to firehose
	seenMessageChannel <- m
}

// BotNameFromContext returns the botname from a provided context
func BotNameFromContext(ctx context.Context) string {
	botname, ok := ctx.Value(ServerBotNameContextKey).(string)
	if !ok {
		return defaultBotName
	}
	return botname
}

// BotIDFromContext returns the bot userid from a provided context
func BotIDFromContext(ctx context.Context) string {
	botname, ok := ctx.Value(ServerBotIDContextKey).(string)
	if !ok {
		return defaultBotID
	}
	return botname
}

// generate a full rtminfo response for initial rtm connections
func generateRTMInfo(ctx context.Context, wsurl string) *fullInfoSlackResponse {
	rtmInfo := slack.Info{
		URL:  wsurl,
		Team: defaultTeam,
		User: defaultBotInfo,
	}
	rtmInfo.User.ID = BotIDFromContext(ctx)
	rtmInfo.User.Name = BotNameFromContext(ctx)
	return &fullInfoSlackResponse{
		rtmInfo,
		okWebResponse,
	}
}

func nowAsJSONTime() slack.JSONTime {
	return slack.JSONTime(time.Now().Unix())
}

func defaultBotInfoJSON(ctx context.Context) string {
	botid := BotIDFromContext(ctx)
	botname := BotNameFromContext(ctx)
	return fmt.Sprintf(`
		{
			"ok":true,
			"bot":{
					"id": "%s",
					"app_id": "A4H1JB4AZ",
					"deleted": false,
					"name": "%s",
					"icons": {
						"image_36": "https://localhost.localdomain/img36.png",
						"image_48": "https://localhost.localdomain/img48.png",
						"image_72": "https://localhost.localdomain/img72.png"
					}
				}
			}
		}
		`, botid, botname)
}
