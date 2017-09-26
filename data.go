package slacktest

import (
	"time"

	slack "github.com/nlopes/slack"
)

const defaultBotName = "TestSlackBot"
const defaultBotID = "U023BECGF"
const defaultTeamID = "T024BE7LD"
const defaultTeamName = "SlackTest Team"
const defaultTeamDomain = "testdomain"
const defaultTeamEmailDomain = "testdomain.local"

var defaultCreatedTs = slack.JSONTime(time.Now().Unix())

var defaultTeam = &slack.Team{
	ID:     defaultTeamID,
	Name:   defaultTeamName,
	Domain: defaultTeamDomain,
}

var defaultBotInfo = &slack.UserDetails{
	ID:             defaultBotID,
	Name:           defaultBotName,
	Created:        defaultCreatedTs,
	ManualPresence: "true",
	Prefs:          slack.UserPrefs{},
}

var okWebResponse = slack.WebResponse{
	Ok: true,
}

var errWebResponse = slack.WebResponse{
	Ok: false,
}
