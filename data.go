package slacktest

import (
	"fmt"

	slack "github.com/nlopes/slack"
)

const defaultBotName = "TestSlackBot"
const defaultBotID = "U023BECGF"
const defaultTeamID = "T024BE7LD"
const defaultTeamName = "SlackTest Team"
const defaultTeamDomain = "testdomain"

var defaultCreatedTs = nowAsJSONTime()

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

var defaultChannelsListJSON = fmt.Sprintf(`
	{
		"ok": true,
		"channels": [%s, %s]
	}
	`, defaultGeneralChannelJSON, defaultExtraChannelJSON)

var defaultGroupsListJSON = fmt.Sprintf(`
		{
			"ok": true,
			"groups": [%s]
		}
		`, defaultGroupJSON)

var defaultUsersInfoJSON = fmt.Sprintf(`
	{
		"ok":true,
		%s
	}
	`, defaultNonBotUser)

var defaultGeneralChannelJSON = fmt.Sprintf(`
	{
        "id": "C024BE91L",
        "name": "general",
        "is_channel": true,
        "created": %d,
        "creator": "W012A3CDE",
        "is_archived": false,
        "is_general": true,

        "members": [
            "W012A3CDE"
        ],

        "topic": {
            "value": "Fun times",
            "creator": "W012A3CDE",
            "last_set": %d
        },
        "purpose": {
            "value": "This channel is for fun",
            "creator": "W012A3CDE",
            "last_set": %d
        },

        "is_member": true
    }
`, nowAsJSONTime(), nowAsJSONTime(), nowAsJSONTime())

var defaultExtraChannelJSON = fmt.Sprintf(`
	{
        "id": "C024BE92L",
        "name": "bot-playground",
        "is_channel": true,
        "created": %d,
        "creator": "W012A3CDE",
        "is_archived": false,
        "is_general": true,

        "members": [
            "W012A3CDE"
        ],

        "topic": {
            "value": "Fun times",
            "creator": "W012A3CDE",
            "last_set": %d
        },
        "purpose": {
            "value": "This channel is for fun",
            "creator": "W012A3CDE",
            "last_set": %d
        },

        "is_member": true
    }
`, nowAsJSONTime(), nowAsJSONTime(), nowAsJSONTime())

var defaultGroupJSON = fmt.Sprintf(`{
    "id": "G024BE91L",
    "name": "secretplans",
    "is_group": true,
    "created": %d,
    "creator": "W012A3CDE",
    "is_archived": false,
    "members": [
        "W012A3CDE"
    ],
    "topic": {
        "value": "Secret plans on hold",
        "creator": "W012A3CDE",
        "last_set": %d
    },
    "purpose": {
        "value": "Discuss secret plans that no-one else should know",
        "creator": "W012A3CDE",
        "last_set": %d
    }
}`, nowAsJSONTime(), nowAsJSONTime(), nowAsJSONTime())

var defaultNonBotUser = `
		"user": {
			"id": "W012A3CDE",
			"team_id": "T012AB3C4",
			"name": "spengler",
			"deleted": false,
			"color": "9f69e7",
			"real_name": "Egon Spengler",
			"tz": "America/Los_Angeles",
			"tz_label": "Pacific Daylight Time",
			"tz_offset": -25200,
			"profile": {
				"avatar_hash": "ge3b51ca72de",
				"status_text": "Print is dead",
				"status_emoji": ":books:",
				"real_name": "Egon Spengler",
				"display_name": "spengler",
				"real_name_normalized": "Egon Spengler",
				"display_name_normalized": "spengler",
				"email": "spengler@ghostbusters.example.com",
				"image_24": "https://localhost.localdomain/avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
				"image_32": "https://localhost.localdomain/avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
				"image_48": "https://localhost.localdomain/avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
				"image_72": "https://localhost.localdomain/avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
				"image_192": "https://localhost.localdomain/avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
				"image_512": "https://localhost.localdomain/avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
				"team": "T012AB3C4"
			},
			"is_admin": true,
			"is_owner": false,
			"is_primary_owner": false,
			"is_restricted": false,
			"is_ultra_restricted": false,
			"is_bot": false,
			"is_stranger": false,
			"updated": 1502138686,
			"is_app_user": false,
			"has_2fa": false,
			"locale": "en-US"
		}
	`
