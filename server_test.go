package slacktest

import (
	"encoding/json"
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
	s.SendMessageToChannel("C123456789", t.Name())
	time.Sleep(2 * time.Second)
	assert.True(t, s.SawOutgoingMessage(t.Name()))
	s.Stop()
}

func TestServerSendMessageToBot(t *testing.T) {
	s := NewTestServer()
	go s.Start()
	s.SendMessageToBot("C123456789", t.Name())
	expectedMsg := fmt.Sprintf("<@%s> %s", s.BotID, t.Name())
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

func TestBotDirectMessage(t *testing.T) {
	s := NewTestServer()
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	s.SendDirectMessageToBot(t.Name())
	expectedMsg := fmt.Sprintf(t.Name())
	time.Sleep(2)
	assert.True(t, s.SawOutgoingMessage(expectedMsg))
	s.Stop()
}

func TestGetSeenOutboundMessages(t *testing.T) {
	maxWait := 5 * time.Second
	s := NewTestServer()
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	s.SendMessageToChannel("foo", "should see this message")
	time.Sleep(maxWait)
	seenOutbound := s.GetSeenOutboundMessages()
	assert.True(t, len(seenOutbound) > 0)
	hadMessage := false
	for _, msg := range seenOutbound {
		var m = slack.Message{}
		jerr := json.Unmarshal([]byte(msg), &m)
		assert.NoError(t, jerr, "messages should decode as slack.Message")
		if m.Text == "should see this message" {
			hadMessage = true
			break
		}
	}
	assert.True(t, hadMessage, "did not see my sent message")
}

func TestGetSeenInboundMessages(t *testing.T) {
	maxWait := 5 * time.Second
	s := NewTestServer()
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	api := slack.New("ABCDEFG")
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	rtm.SendMessage(&slack.OutgoingMessage{
		Channel: "foo",
		Text:    "should see this inbound message",
	})
	time.Sleep(maxWait)
	seenInbound := s.GetSeenInboundMessages()
	assert.True(t, len(seenInbound) > 0)
	hadMessage := false
	for _, msg := range seenInbound {
		var m = slack.Message{}
		jerr := json.Unmarshal([]byte(msg), &m)
		assert.NoError(t, jerr, "messages should decode as slack.Message")
		if m.Text == "should see this inbound message" {
			hadMessage = true
			break
		}
	}
	assert.True(t, hadMessage, "did not see my sent message")
	assert.True(t, s.SawMessage("should see this inbound message"))
}

func TestSendToWebsocket(t *testing.T) {
	maxWait := 5 * time.Second
	s := NewTestServer()
	go s.Start()
	_, rtm := s.GetTestRTMInstance()
	go rtm.ManageConnection()
	evChan := make(chan (slack.Channel), 1)
	go func() {
		for msg := range rtm.IncomingEvents {
			switch ev := msg.Data.(type) {
			case *slack.ChannelJoinedEvent:
				evChan <- ev.Channel
			}
		}
	}()
	e := fmt.Sprintf(`{"type":"channel_joined", "channel":%s}`, defaultExtraChannelJSON)
	s.SendToWebsocket(e)
	time.Sleep(maxWait)
	select {
	case m := <-evChan:
		assert.Equal(t, "C024BE92L", m.ID, "channel id should match")
		assert.Equal(t, "Fun times", m.Topic.Value, "topic should match")
		s.Stop()
		break
	case <-time.After(maxWait):
		assert.FailNow(t, "did not get channel joined event in time")
	}

}
