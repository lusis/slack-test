package slacktest

import (
	"context"

	slackbot "github.com/lusis/go-slackbot"
	slack "github.com/nlopes/slack"
)

func testSlackBotEchoHandler(ctx context.Context, b *slackbot.Bot, evt *slack.MessageEvent) {
	b.Reply(evt, "bot saw: "+evt.Text, slackbot.WithoutTyping)
}
