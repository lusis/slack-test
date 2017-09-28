package slacktest

import (
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
	channels := &serverChannels{}
	groups := &serverGroups{}
	s := &Server{}
	mux := http.NewServeMux()
	mux.Handle("/ws", contextHandler(s, wsHandler))
	mux.Handle("/rtm.start", contextHandler(s, rtmStartHandler))
	mux.Handle("/chat.postMessage", contextHandler(s, postMessageHandler))
	mux.Handle("/channels.list", contextHandler(s, listChannelsHandler))
	mux.Handle("/groups.list", contextHandler(s, listGroupsHandler))
	mux.Handle("/users.info", contextHandler(s, usersInfoHandler))
	mux.Handle("/bots.info", contextHandler(s, botsInfoHandler))
	httpserver := httptest.NewUnstartedServer(mux)
	addr := httpserver.Listener.Addr().String()

	s.ServerAddr = addr
	s.server = httpserver
	s.BotName = defaultBotName
	s.BotID = defaultBotID
	s.SeenFeed = seenMessageChannel
	s.channels = channels
	s.groups = groups
	return s
}

// GetChannels returns all the fake channels registered
func (sts *Server) GetChannels() []slack.Channel {
	sts.channels.RLock()
	defer sts.channels.RUnlock()
	return sts.channels.channels
}

// GetGroups returns all the fake groups registered
func (sts *Server) GetGroups() []slack.Group {
	return sts.groups.channels
}

// AddChannel adds a new fake channel
func (sts *Server) AddChannel(c slack.Channel) {
	sts.channels.Lock()
	sts.channels.channels = append(sts.channels.channels, c)
	sts.channels.Unlock()
}

// AddGroup adds a new fake group
func (sts *Server) AddGroup(c slack.Group) {
	sts.groups.Lock()
	sts.groups.channels = append(sts.groups.channels, c)
	sts.groups.Unlock()
}

// GetSeenMessages returns all messages seen via websocket excluding pings
func (sts *Server) GetSeenMessages() []string {
	seenInboundMessages.RLock()
	m := seenInboundMessages.messages
	seenInboundMessages.RUnlock()
	return m
}

// SawOutgoingMessage checks if a message was sent to connected websocket clients
func (sts *Server) SawOutgoingMessage(msg string) bool {
	seenOutboundMessages.RLock()
	defer seenOutboundMessages.RUnlock()
	for _, m := range seenOutboundMessages.messages {
		evt := &slack.MessageEvent{}
		jErr := json.Unmarshal([]byte(m), evt)
		if jErr != nil {
			continue
		}
		if evt.Text == msg {
			return true
		}
	}
	return false
}

// SawMessage checks if an incoming message was seen
func (sts *Server) SawMessage(msg string) bool {
	seenInboundMessages.RLock()
	defer seenInboundMessages.RUnlock()
	for _, m := range seenInboundMessages.messages {
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
	m.Type = slack.TYPE_MESSAGE
	m.Channel = channel
	m.User = "W012A3CDE"
	m.Text = fmt.Sprintf("<@%s> %s", sts.BotID, msg)
	m.Timestamp = fmt.Sprintf("%d", time.Now().Unix())
	j, jErr := json.Marshal(m)
	if jErr != nil {
		log.Printf("Unable to marshal message for bot: %s", jErr.Error())
		return
	}
	go queueForWebsocket(string(j))
}

// SendMessageToChannel sends a message to a channel
func (sts *Server) SendMessageToChannel(channel, msg string) {
	m := slack.Message{}
	m.Type = slack.TYPE_MESSAGE
	m.Channel = channel
	m.Text = msg
	m.User = "W012A3CDE"
	m.Timestamp = fmt.Sprintf("%d", time.Now().Unix())
	j, jErr := json.Marshal(m)
	if jErr != nil {
		log.Printf("Unable to marshal message for channel: %s", jErr.Error())
		return
	}
	stringMsg := string(j)
	go queueForWebsocket(stringMsg)
}

// SetBotName sets a custom botname
func (sts *Server) SetBotName(b string) {
	sts.BotName = b
}
