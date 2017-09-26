package slacktest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	slack "github.com/nlopes/slack"
)

// NewTestServer returns a slacktest.Server ready to be started
func NewTestServer() *Server {
	sendMessageChannel = make(chan (string))
	seenMessageChannel = make(chan (string))
	s := &Server{
		SendMessages: sendMessageChannel,
		SeenFeed:     seenMessageChannel,
	}
	mux := http.NewServeMux()
	mux.Handle("/ws", contextHandler(s, wsHandler))
	mux.Handle("/rtm.start", contextHandler(s, rtmStartHandler))
	mux.Handle("/chat.postMessage", contextHandler(s, postMessageHandler))
	httpserver := httptest.NewUnstartedServer(mux)
	addr := httpserver.Listener.Addr().String()

	s.ServerAddr = addr
	s.server = httpserver
	s.BotName = defaultBotName
	s.BotID = defaultBotID
	return s
}

// GetSeenMessages returns all messages seen via websocket excluding pings
func (sts *Server) GetSeenMessages() []string {
	return seenMessages
}

// SawMessage checks if a message was seen
func (sts *Server) SawMessage(msg string) bool {
	for _, m := range seenMessages {
		evt := &slack.MessageEvent{}
		jErr := json.Unmarshal([]byte(m), evt)
		if jErr != nil {
			// This event isn't a message event so we'll skip it
			continue
		}
		if evt.Text == msg {
			return true
		}
	}
	return false
}

// GetAPIURL returns the api url you can pass to slack.SLACK_API
func (sts *Server) GetAPIURL() string {
	return "http://" + sts.ServerAddr + "/"
}

// GetWSURL returns the websocket url
func (sts *Server) GetWSURL() string {
	return "ws://" + sts.ServerAddr + "/ws"
}

// Stop stops the test server
func (sts *Server) Stop() {
	sts.server.Close()
}

// Start starts the test server
func (sts *Server) Start() {
	sts.server.Start()
}

// SendMessageToBot sends a message addressed to the Bot
func (sts *Server) SendMessageToBot(channel, msg string) {
	m := slack.Message{}
	m.Type = "message"
	m.Channel = channel
	m.Text = fmt.Sprintf("<@%s> %s", sts.BotID, msg)
	m.Timestamp = fmt.Sprintf("%d", time.Now().Unix())
	j, jErr := json.Marshal(m)
	if jErr != nil {
		log.Printf("Unable to marshal message for bot: %s", jErr.Error())
		return
	}
	sts.SendMessages <- string(j)
}

// SendMessageToChannel sends a message to a channel
func (sts *Server) SendMessageToChannel(channel, msg string) {
	m := slack.Message{}
	m.Type = "message"
	m.Channel = channel
	m.Text = msg
	m.Timestamp = fmt.Sprintf("%d", time.Now().Unix())
	j, jErr := json.Marshal(m)
	if jErr != nil {
		log.Printf("Unable to marshal message for channel: %s", jErr.Error())
		return
	}
	sts.SendMessages <- string(j)
}

// SetBotName sets a custom botname
func (sts *Server) SetBotName(b string) {
	sts.BotName = b
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
