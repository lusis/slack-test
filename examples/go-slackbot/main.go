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

func directMessageHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "sorry I can't do direct messages", slackbot.WithoutTyping)
}

func configureBot(bot *slackbot.Bot) {
	bot.Hear("global message").MessageHandler(globalMessageHandler)
	toMe := bot.Messages(slackbot.DirectMention).Subrouter()
	toMe.Hear("greetings and salutations").MessageHandler(helloFunc)
	dms := bot.Messages(slackbot.DirectMessage).Subrouter()
	dms.Hear("^.*$").MessageHandler(directMessageHandler)
}
func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))
	configureBot(bot)
	bot.Run()
}
