package slacktest

import (
	"fmt"
	"testing"
	"time"

	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

func TestDefaultNewServer(t *testing.T) {
	s := NewTestServer()
	assert.Equal(t, defaultBotID, s.BotID)
	assert.Equal(t, defaultBotName, s.BotName)
	assert.NotEmpty(t, s.ServerAddr)
	s.Stop()
}

func TestCustomNewServer(t *testing.T) {
	s := NewTestServer()
	s.SetBotName("BobsBot")
	assert.Equal(t, "BobsBot", s.BotName)
}

func TestServerSendMessageToChannel(t *testing.T) {
	s := NewTestServer()
	go s.Start()
	s.SendMessageToChannel("C123456789", "test message")
	time.Sleep(2 * time.Second)
	assert.True(t, s.SawOutgoingMessage("test message"))
	s.Stop()
}

func TestServerSendMessageToBot(t *testing.T) {
	s := NewTestServer()
	go s.Start()
	s.SendMessageToBot("C123456789", "bot message")
	expectedMsg := fmt.Sprintf("<@%s> bot message", s.BotID)
	time.Sleep(2 * time.Second)
	assert.True(t, s.SawOutgoingMessage(expectedMsg))
	s.Stop()
}

func TestServerListChannels(t *testing.T) {
	s := NewTestServer()
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	client := slack.New("ABCDEFG")
	channels, err := client.GetChannels(true)
	assert.NoError(t, err)
	assert.Len(t, channels, 2)
	assert.Equal(t, "C024BE91L", channels[0].ID)
	assert.Equal(t, "C024BE92L", channels[1].ID)
	for _, channel := range channels {
		assert.Equal(t, "W012A3CDE", channel.Creator)
	}
}

func TestBotsInfoHandler(t *testing.T) {
	s := NewTestServer()
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	client := slack.New("ABCDEFG")
	user, err := client.GetUserInfo("123456")
	assert.NoError(t, err)
	assert.Equal(t, "W012A3CDE", user.ID)
	assert.Equal(t, "spengler", user.Name)
	assert.True(t, user.IsAdmin)
}
