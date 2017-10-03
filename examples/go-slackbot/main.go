package main

import (
	"context"
	"fmt"
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

func postMessageHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	_, _, err := bot.Client.PostMessage(evt.Channel, "posting to a channel via api", slack.PostMessageParameters{AsUser: true})
	if err != nil {
		fmt.Printf("Got an error making an api call to post a message: %s\n", err.Error())
	}
}

func withAttachmentHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	p := &slack.PostMessageParameters{
		AsUser: true,
		Attachments: []slack.Attachment{
			slack.Attachment{
				Fallback:   "this is the fallback text",
				AuthorName: "Message Author Name",
				Title:      "message title",
				Text:       "message text",
				Fields: []slack.AttachmentField{
					slack.AttachmentField{
						Title: "field title",
						Value: "field value",
						Short: true,
					},
				},
			},
		},
	}
	_, _, err := bot.Client.PostMessage(evt.Channel, "", *p)
	if err != nil {
		fmt.Printf("got an error making an api call to post a message with an attachment: %s\n", err.Error())
	}
}

func channelJoinHandler(ctx context.Context, bot *slackbot.Bot, channel *slack.Channel) {
	_, _, err := bot.Client.PostMessage(channel.ID, "thanks for the invite", slack.PostMessageParameters{})
	if err != nil {
		fmt.Printf("error handling channel join event: %s", err.Error())
		return
	}
}

func configureBot(bot *slackbot.Bot) {
	bot.OnChannelJoin(channelJoinHandler)
	bot.Hear("send an attachment").MessageHandler(withAttachmentHandler)
	bot.Hear("send to api").MessageHandler(postMessageHandler)
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
