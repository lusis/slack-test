package slacktest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateDefaultRTMInfo(t *testing.T) {
	wsurl := "ws://127.0.0.1:5555/ws"
	ctx := context.TODO()
	info := generateRTMInfo(ctx, wsurl)
	require.Equal(t, wsurl, info.URL)
	require.True(t, info.Ok)
	require.Equal(t, defaultBotID, info.User.ID)
	require.Equal(t, defaultBotName, info.User.Name)
	require.Equal(t, defaultTeamID, info.Team.ID)
	require.Equal(t, defaultTeamName, info.Team.Name)
	require.Equal(t, defaultTeamDomain, info.Team.Domain)
}

func TestCustomDefaultRTMInfo(t *testing.T) {
	wsurl := "ws://127.0.0.1:5555/ws"
	ctx := context.TODO()
	ctx = context.WithValue(ctx, ServerBotIDContextKey, "U1234567890")
	ctx = context.WithValue(ctx, ServerBotNameContextKey, "SomeTestBotThing")
	info := generateRTMInfo(ctx, wsurl)
	require.Equal(t, wsurl, info.URL)
	require.True(t, info.Ok)
	require.Equal(t, "U1234567890", info.User.ID)
	require.Equal(t, "SomeTestBotThing", info.User.Name)
	require.Equal(t, defaultTeamID, info.Team.ID)
	require.Equal(t, defaultTeamName, info.Team.Name)
	require.Equal(t, defaultTeamDomain, info.Team.Domain)
}

func TestGetHubMissingServerAddr(t *testing.T) {
	mc, err := getHubForServer("")
	require.Nil(t, mc.seen, "seen should be nil")
	require.Nil(t, mc.sent, "sent should be nil")
	require.Nil(t, mc.posted, "posted should be nil")
	require.Error(t, err, "should return an error")
	require.EqualError(t, err, ErrPassedEmptyServerAddr.Error())
}

func TestGetHubNoQueuesForServer(t *testing.T) {
	mc, err := getHubForServer("foo")
	require.Nil(t, mc.seen, "seen should be nil")
	require.Nil(t, mc.sent, "sent should be nil")
	require.Nil(t, mc.posted, "posted should be nil")
	require.Error(t, err, "should return an error")
	require.EqualError(t, err, ErrNoQueuesRegisteredForServer.Error())
}

func TestUnableToAddToHub(t *testing.T) {
	err := addServerToHub(&Server{}, &messageChannels{})
	require.Error(t, err, "should return and error")
	require.EqualError(t, err, ErrEmptyServerToHub.Error())
}
