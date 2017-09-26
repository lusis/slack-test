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
	slack.SLACK_API = "http://" + s.ServerAddr + "/"
	go s.Start()
	bot := slackbot.New("ABCDEFG")
	bot.Hear("global message").MessageHandler(globalMessageHandler)
	go bot.Run()
	s.SendMessageToChannel("#general", "global message")
	time.Sleep(2 * time.Second)
	assert.True(t, s.SawMessage("I see your global message"), "bot did not respond correctly")
	s.Stop()
}

func TestHelloMessageHandler(t *testing.T) {
	s := slacktest.NewTestServer()
	s.SetBotName("foobot")
	slack.SLACK_API = "http://" + s.ServerAddr + "/"
	go s.Start()
	bot := slackbot.New("ABCDEFG")
	toMe := bot.Messages(slackbot.DirectMention).Subrouter()
	toMe.Hear("greetings and salutations").MessageHandler(helloFunc)
	go bot.Run()
	s.SendMessageToBot("#general", "greetings and salutations")
	time.Sleep(2 * time.Second)
	assert.True(t, s.SawMessage("hi there to you too!"), "bot did not respond correctly")
	s.Stop()
}
