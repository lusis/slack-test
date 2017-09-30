package main

import (
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
	slack.SLACK_API = s.GetAPIURL()
	go s.Start()
	bot := slackbot.New("ABCDEFG")
	configureBot(bot)
	go bot.Run()
	s.SendMessageToChannel("#general", "global message")
	time.Sleep(2 * time.Second)
	assert.True(t, s.SawMessage("I see your global message"), "bot did not respond correctly")
	s.Stop()
}

func TestHelloMessageHandler(t *testing.T) {
	s := slacktest.NewTestServer()
	s.SetBotName("foobot")
	slack.SLACK_API = s.GetAPIURL()
	go s.Start()
	bot := slackbot.New("ABCDEFG")
	configureBot(bot)
	go bot.Run()
	s.SendMessageToBot("#general", "greetings and salutations")
	time.Sleep(2 * time.Second)
	assert.True(t, s.SawMessage("hi there to you too!"), "bot did not respond correctly")
	s.Stop()
}

func TestDirectMessageHandler(t *testing.T) {
	s := slacktest.NewTestServer()
	s.SetBotName("foobot")
	slack.SLACK_API = s.GetAPIURL()
	go s.Start()
	bot := slackbot.New("ABCDEFG")
	configureBot(bot)
	go bot.Run()
	s.SendDirectMessageToBot("wanna chat?")
	time.Sleep(2 * time.Second)
	assert.True(t, s.SawMessage("sorry I can't do direct messages"), "bot did not respond correctly")
	s.Stop()
}
