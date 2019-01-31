package main

import (
	"encoding/json"
	"testing"
	"time"

	slackbot "github.com/lusis/go-slackbot"
	slacktest "github.com/lusis/slack-test"
	slack "github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

func TestGlobalMessageHandler(t *testing.T) {
	s := slacktest.NewTestServer()
	s.SetBotName("TestSlackBot")
	slack.APIURL = s.GetAPIURL()
	go s.Start()
	bot := slackbot.New("ABCDEFG")
	configureBot(bot)
	go bot.Run()
	s.SendMessageToChannel("C024BE91L", "global message")
	time.Sleep(2 * time.Second)
	assert.True(t, s.SawMessage("I see your global message"), "bot did not respond correctly")
	s.Stop()
}

func TestHelloMessageHandler(t *testing.T) {
	s := slacktest.NewTestServer()
	s.SetBotName("foobot")
	slack.APIURL = s.GetAPIURL()
	go s.Start()
	bot := slackbot.New("ABCDEFG")
	configureBot(bot)
	go bot.Run()
	s.SendMessageToBot("#C024BE91L", "greetings and salutations")
	time.Sleep(2 * time.Second)
	assert.True(t, s.SawMessage("hi there to you too!"), "bot did not respond correctly")
	s.Stop()
}

func TestDirectMessageHandler(t *testing.T) {
	s := slacktest.NewTestServer()
	s.SetBotName("foobot")
	slack.APIURL = s.GetAPIURL()
	go s.Start()
	bot := slackbot.New("ABCDEFG")
	configureBot(bot)
	go bot.Run()
	s.SendDirectMessageToBot("wanna chat?")
	time.Sleep(2 * time.Second)
	assert.True(t, s.SawMessage("sorry I can't do direct messages"), "bot did not respond correctly")
	s.Stop()
}

func TestPostMessageHandler(t *testing.T) {
	s := slacktest.NewTestServer()
	slack.APIURL = s.GetAPIURL()
	go s.Start()
	bot := slackbot.New("ABCDEFG")
	configureBot(bot)
	//bot.Client.SetDebug(true)
	go bot.Run()
	s.SendMessageToChannel("C024BE91L", "send to api")
	time.Sleep(2 * time.Second)
	seenMessages := s.GetSeenOutboundMessages()
	if !assert.Len(t, seenMessages, 2, "should only have two messages") {
		t.FailNow()
	}
	var m = slack.Message{}
	jErr := json.Unmarshal([]byte(seenMessages[1]), &m)
	assert.NoError(t, jErr, "message should decode properly")
	assert.Equal(t, s.BotName, m.Username)
}

func TestPostAttachmentHandler(t *testing.T) {
	s := slacktest.NewTestServer()
	slack.APIURL = s.GetAPIURL()
	go s.Start()
	bot := slackbot.New("ABCDEFG")
	configureBot(bot)
	//bot.Client.SetDebug(true)
	go bot.Run()
	s.SendMessageToChannel("C024BE91L", "send an attachment")
	time.Sleep(2 * time.Second)
	seenMessages := s.GetSeenOutboundMessages()
	if !assert.Len(t, seenMessages, 2, "should only have two messages") {
		t.FailNow()
	}
	var m = slack.Message{}
	jErr := json.Unmarshal([]byte(seenMessages[1]), &m)
	assert.NoError(t, jErr, "message should decode properly")
	assert.Len(t, m.Attachments, 1, "message should have one attachment")
	assert.Len(t, m.Attachments[0].Fields, 1, "message should have one field")
	assert.Equal(t, m.Attachments[0].Fields[0].Title, "field title", "field should have the correct title")
}

func TestChannelJoinHandler(t *testing.T) {
	s := slacktest.NewTestServer()
	slack.APIURL = s.GetAPIURL()
	go s.Start()
	bot := slackbot.New("ABCDEFG")
	configureBot(bot)
	//bot.Client.SetDebug(true)
	go bot.Run()
	s.SendBotChannelInvite()
	time.Sleep(2 * time.Second)
	seenMessages := s.GetSeenOutboundMessages()
	if !assert.Len(t, seenMessages, 2, "should only have two messages") {
		t.FailNow()
	}
	var m = slack.Message{}
	jErr := json.Unmarshal([]byte(seenMessages[1]), &m)
	if !assert.NoError(t, jErr, "message should decode properly") {
		t.FailNow()
	}
	assert.Equal(t, m.Text, "thanks for the invite", "bot should send message on invite")
	assert.Equal(t, m.Channel, "C024BE92L", "message should be to invited channel")
}

func TestChannelJoinHandlerGroup(t *testing.T) {
	s := slacktest.NewTestServer()
	slack.APIURL = s.GetAPIURL()
	go s.Start()
	bot := slackbot.New("ABCDEFG")
	configureBot(bot)
	//bot.Client.SetDebug(true)
	go bot.Run()
	s.SendBotGroupInvite()
	time.Sleep(2 * time.Second)
	seenMessages := s.GetSeenOutboundMessages()
	if !assert.Len(t, seenMessages, 2, "should only have two messages") {
		t.FailNow()
	}
	var m = slack.Message{}
	jErr := json.Unmarshal([]byte(seenMessages[1]), &m)
	if !assert.NoError(t, jErr, "message should decode properly") {
		t.FailNow()
	}
	assert.Equal(t, m.Text, "thanks for the invite", "bot should send message on invite")
	assert.Equal(t, m.Channel, "G024BE91L", "message should be to invited group")
}
