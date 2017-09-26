package main

import (
	"context"
	"os"

	slackbot "github.com/lusis/go-slackbot"
	slack "github.com/nlopes/slack"
)

func helloFunc(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "hi there to you too!", slackbot.WithoutTyping)
}

func globalMessageHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "I see your global message", slackbot.WithoutTyping)
}

func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))
	bot.Hear("global message").MessageHandler(globalMessageHandler)
	toMe := bot.Messages(slackbot.DirectMention).Subrouter()
	toMe.Hear("greetings and salutations").MessageHandler(helloFunc)
	bot.Run()
}
