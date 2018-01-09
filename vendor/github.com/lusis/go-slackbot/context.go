package slackbot

import (
	"context"

	"github.com/nlopes/slack"
)

const (
	// BotContext is the context key for the bot context entry
	BotContext = "__BOT_CONTEXT__"
	// MessageContext is the context key for the message context entry
	MessageContext = "__MESSAGE_CONTEXT__"
	// NamedCaptureContextKey is the key for named captures
	NamedCaptureContextKey = "__NAMED_CAPTURES__"
)

// BotFromContext creates a Bot from provided Context
func BotFromContext(ctx context.Context) *Bot {
	if result, ok := ctx.Value(contextKey(BotContext)).(*Bot); ok {
		return result
	}
	return nil
}

// AddBotToContext sets the bot reference in context and returns the newly derived context
func AddBotToContext(ctx context.Context, bot *Bot) context.Context {
	return context.WithValue(ctx, contextKey(BotContext), bot)
}

// MessageFromContext gets the message from the provided context
func MessageFromContext(ctx context.Context) *slack.MessageEvent {
	if result, ok := ctx.Value(contextKey(MessageContext)).(*slack.MessageEvent); ok {
		return result
	}
	return nil
}

// AddMessageToContext sets the Slack message event reference in context and returns the newly derived context
func AddMessageToContext(ctx context.Context, msg *slack.MessageEvent) context.Context {
	return context.WithValue(ctx, contextKey(MessageContext), msg)
}

// NamedCapturesFromContext returns any NamedCaptures parsed from regexp
func NamedCapturesFromContext(ctx context.Context) NamedCaptures {
	if result, ok := ctx.Value(contextKey(NamedCaptureContextKey)).(NamedCaptures); ok {
		return result
	}
	return NamedCaptures{}
}
