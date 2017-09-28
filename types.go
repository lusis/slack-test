package slacktest

import (
	"log"
	"net/http"
	"net/http/httptest"
	"sync"

	slack "github.com/nlopes/slack"
)

type contextKey string

// ServerURLContextKey is the context key to store the server's url
var ServerURLContextKey contextKey = "__SERVER_URL__"

// ServerWSContextKey is the context key to store the server's ws url
var ServerWSContextKey contextKey = "__SERVER_WS_URL__"

// ServerBotNameContextKey is the bot name
var ServerBotNameContextKey contextKey = "__SERVER_BOTNAME__"

// ServerBotIDContextKey is the bot userid
var ServerBotIDContextKey contextKey = "__SERVER_BOTID__"

// ServerBotChannelsContextKey is the list of channels associated with the fake server
var ServerBotChannelsContextKey contextKey = "__SERVER_CHANNELS__"

// ServerBotGroupsContextKey is the list of channels associated with the fake server
var ServerBotGroupsContextKey contextKey = "__SERVER_GROUPS__"

var sendMessageChannel chan (string)
var seenMessageChannel chan (string)
var seenInboundMessages = &messageCollection{}
var seenOutboundMessages = &messageCollection{}

type messageCollection struct {
	sync.RWMutex
	messages []string
}

type serverChannels struct {
	sync.RWMutex
	channels []slack.Channel
}

type serverGroups struct {
	sync.RWMutex
	channels []slack.Group
}

// Server represents a Slack Test server
type Server struct {
	server     *httptest.Server
	mux        *http.ServeMux
	Logger     *log.Logger
	BotName    string
	BotID      string
	ServerAddr string
	SeenFeed   chan (string)
	channels   *serverChannels
	groups     *serverGroups
}

type fullInfoSlackResponse struct {
	slack.Info
	slack.WebResponse
}
