package slacktest

import (
	"testing"

	slack "github.com/nlopes/slack"
	"github.com/stretchr/testify/require"
)

func TestPostMessageHandler(t *testing.T) {
	s, err := NewTestServer()
	require.NoError(t, err)
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	client := slack.New("ABCDEFG")
	channel, tstamp, err := client.PostMessage("foo", t.Name(), slack.PostMessageParameters{})
	require.NoError(t, err, "should not error out")
	require.Equal(t, "foo", channel, "channel should be correct")
	require.NotEmpty(t, tstamp, "timestamp should not be empty")
}

func TestServerListChannels(t *testing.T) {
	s, err := NewTestServer()
	require.NoError(t, err)
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	client := slack.New("ABCDEFG")
	channels, err := client.GetChannels(true)
	require.NoError(t, err)
	require.Len(t, channels, 2)
	require.Equal(t, "C024BE91L", channels[0].ID)
	require.Equal(t, "C024BE92L", channels[1].ID)
	for _, channel := range channels {
		require.Equal(t, "W012A3CDE", channel.Creator)
	}
}

func TestUserInfoHandler(t *testing.T) {
	s, err := NewTestServer()
	require.NoError(t, err)
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	client := slack.New("ABCDEFG")
	user, err := client.GetUserInfo("123456")
	require.NoError(t, err)
	require.Equal(t, "W012A3CDE", user.ID)
	require.Equal(t, "spengler", user.Name)
	require.True(t, user.IsAdmin)
}

func TestBotInfoHandler(t *testing.T) {
	s, err := NewTestServer()
	require.NoError(t, err)
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	client := slack.New("ABCDEFG")
	bot, err := client.GetBotInfo(s.BotID)
	require.NoError(t, err)
	require.Equal(t, s.BotID, bot.ID)
	require.Equal(t, s.BotName, bot.Name)
	require.False(t, bot.Deleted)
}

func TestListGroupsHandler(t *testing.T) {
	s, err := NewTestServer()
	require.NoError(t, err)
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	client := slack.New("ABCDEFG")
	groups, err := client.GetGroups(true)
	require.NoError(t, err)
	require.Len(t, groups, 1, "should have one group")
	mygroup := groups[0]
	require.Equal(t, "G024BE91L", mygroup.ID, "id should match")
	require.Equal(t, "secretplans", mygroup.Name, "name should match")
	require.True(t, mygroup.IsGroup, "should be a group")
}

func TestListChannelsHandler(t *testing.T) {
	s, err := NewTestServer()
	require.NoError(t, err)
	go s.Start()
	slack.SLACK_API = s.GetAPIURL()
	client := slack.New("ABCDEFG")
	channels, err := client.GetChannels(true)
	require.NoError(t, err)
	require.Len(t, channels, 2, "should have two channels")
	generalChan := channels[0]
	otherChan := channels[1]
	require.Equal(t, "C024BE91L", generalChan.ID, "id should match")
	require.Equal(t, "general", generalChan.Name, "name should match")
	require.Equal(t, "Fun times", generalChan.Topic.Value)
	require.True(t, generalChan.IsMember, "should be in channel")
	require.Equal(t, "C024BE92L", otherChan.ID, "id should match")
	require.Equal(t, "bot-playground", otherChan.Name, "name should match")
	require.Equal(t, "Fun times", otherChan.Topic.Value)
	require.True(t, otherChan.IsMember, "should be in channel")
}
