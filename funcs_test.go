package slacktest

import (
	"context"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGenerateDefaultRTMInfo(t *testing.T) {
	wsurl := "ws://127.0.0.1:5555/ws"
	ctx := context.TODO()
	info := generateRTMInfo(ctx, wsurl)
	assert.Equal(t, wsurl, info.URL)
	assert.True(t, info.Ok)
	assert.Equal(t, defaultBotID, info.User.ID)
	assert.Equal(t, defaultBotName, info.User.Name)
	assert.Equal(t, defaultTeamID, info.Team.ID)
	assert.Equal(t, defaultTeamName, info.Team.Name)
	assert.Equal(t, defaultTeamDomain, info.Team.Domain)
}

func TestCustomDefaultRTMInfo(t *testing.T) {
	wsurl := "ws://127.0.0.1:5555/ws"
	ctx := context.TODO()
	ctx = context.WithValue(ctx, ServerBotIDContextKey, "U1234567890")
	ctx = context.WithValue(ctx, ServerBotNameContextKey, "SomeTestBotThing")
	info := generateRTMInfo(ctx, wsurl)
	assert.Equal(t, wsurl, info.URL)
	assert.True(t, info.Ok)
	assert.Equal(t, "U1234567890", info.User.ID)
	assert.Equal(t, "SomeTestBotThing", info.User.Name)
	assert.Equal(t, defaultTeamID, info.Team.ID)
	assert.Equal(t, defaultTeamName, info.Team.Name)
	assert.Equal(t, defaultTeamDomain, info.Team.Domain)
}
