package slacktest

import (
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert"
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
