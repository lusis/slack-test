package slacktest

import (
	"fmt"

	slack "github.com/nlopes/slack"
)

const defaultBotName = "TestSlackBot"
const defaultBotID = "U023BECGF"
const defaultTeamID = "T024BE7LD"
const defaultNonBotUserID = "W012A3CDE"
const defaultNonBotUserName = "Egon Spengler"
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
        "creator": "%s",
        "is_archived": false,
        "is_general": true,

        "members": [
            "W012A3CDE"
        ],

        "topic": {
            "value": "Fun times",
            "creator": "%s",
            "last_set": %d
        },
        "purpose": {
            "value": "This channel is for fun",
            "creator": "%s",
            "last_set": %d
        },

        "is_member": true
    }
`, nowAsJSONTime(), defaultNonBotUserID, defaultNonBotUserID, nowAsJSONTime(), defaultNonBotUserID, nowAsJSONTime())

var defaultExtraChannelJSON = fmt.Sprintf(`
	{
        "id": "C024BE92L",
        "name": "bot-playground",
        "is_channel": true,
        "created": %d,
        "creator": "%s",
        "is_archived": false,
        "is_general": true,

        "members": [
            "W012A3CDE"
        ],

        "topic": {
            "value": "Fun times",
            "creator": "%s",
            "last_set": %d
        },
        "purpose": {
            "value": "This channel is for fun",
            "creator": "%s",
            "last_set": %d
        },

        "is_member": true
    }
`, nowAsJSONTime(), defaultNonBotUserID, defaultNonBotUserID, nowAsJSONTime(), defaultNonBotUserID, nowAsJSONTime())

var defaultGroupJSON = fmt.Sprintf(`{
    "id": "G024BE91L",
    "name": "secretplans",
    "is_group": true,
    "created": %d,
    "creator": "%s",
    "is_archived": false,
    "members": [
        "W012A3CDE"
    ],
    "topic": {
        "value": "Secret plans on hold",
        "creator": "%s",
        "last_set": %d
    },
    "purpose": {
        "value": "Discuss secret plans that no-one else should know",
        "creator": "%s",
        "last_set": %d
    }
}`, nowAsJSONTime(), defaultNonBotUserID, defaultNonBotUserID, nowAsJSONTime(), defaultNonBotUserID, nowAsJSONTime())

var defaultNonBotUser = fmt.Sprintf(`
		"user": {
			"id": "%s",
			"team_id": "%s",
			"name": "spengler",
			"deleted": false,
			"color": "9f69e7",
			"real_name": "%s",
			"tz": "America/Los_Angeles",
			"tz_label": "Pacific Daylight Time",
			"tz_offset": -25200,
			"profile": {
				"avatar_hash": "ge3b51ca72de",
				"status_text": "Print is dead",
				"status_emoji": ":books:",
				"real_name": "%s",
				"display_name": "spengler",
				"real_name_normalized": "%s",
				"display_name_normalized": "spengler",
				"email": "spengler@ghostbusters.example.com",
				"image_24": "https://localhost.localdomain/avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
				"image_32": "https://localhost.localdomain/avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
				"image_48": "https://localhost.localdomain/avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
				"image_72": "https://localhost.localdomain/avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
				"image_192": "https://localhost.localdomain/avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
				"image_512": "https://localhost.localdomain/avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
				"team": "%s"
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
	`, defaultNonBotUserID, defaultTeamID, defaultNonBotUserName, defaultNonBotUserName, defaultNonBotUserName, defaultTeamID)

var defaultConversationHistoryJSON = fmt.Sprintf(`
	{
        "ok": true,
        "latest": "1522942333",
        "oldest": "1522939726.000713",
        "messages": [
          {
                "type": "message",
                "user": "%s",
                "text": "this has replies",
                "thread_ts": "1522941680.000626",
                "reply_count": 2,
                "replies": [
                  {
                        "user": "%s",
                        "ts": "1522941699.000203"
                  },
                  {
                        "user": "%s",
                        "ts": "1522941709.000423"
                  }
                ],
                "subscribed": false,
                "unread_count": 2,
                "ts": "1522941680.000626"
          },
          {
                "type": "message",
                "user": "%s",
                "text": "this has reactions",
                "ts": "1522940379.000820",
                "reactions": [
                  {
                        "name": "salute",
                        "users": [
                          "U1234"
                        ],
                        "count": 1
                  }
                ]
          },
          {
                "type": "message",
                "subtype": "file_share",
                "text": "<@U1234> uploaded a file: <http://file/-.txt|Untitled>",
                "file": {
                  "id": "F1234",
                  "created": 1522939852,
                  "timestamp": 1522939852,
                  "name": "-.sh",
                  "title": "Untitled",
                  "mimetype": "text/plain",
                  "filetype": "shell",
                  "pretty_type": "Shell",
                  "user": "U1234",
                  "editable": true,
                  "size": 473,
                  "mode": "snippet",
                  "is_external": false,
                  "external_type": "",
                  "is_public": true,
                  "public_url_shared": false,
                  "display_as_bot": false,
                  "username": "",
                  "url_private": "https://file/-.txt",
                  "url_private_download": "http://files/download/-.txt",
                  "permalink": "http://files/files/U1234/F1234/-.txt",
                  "permalink_public": "http://files/1",
                  "edit_link": "http://files/files/U1234/F1233/-.txt/edit",
                  "preview": "<long string>",
                  "preview_highlight": "<preview highlight>",
                  "lines": 1,
                  "lines_more": 0,
                  "preview_is_truncated": false,
                  "channels": [
                        "C6KDDU879"
                  ],
                  "groups": [],
                  "ims": [],
                  "comments_count": 0
                },
                "user": "U753E0PEV",
                "upload": true,
                "display_as_bot": false,
                "username": "bob.smith",
                "bot_id": null,
                "ts": "1522939853.000323"
          },
          {
                "type": "message",
                "user": "%s",
                "text": "",
                "bot_id": "%s",
                "attachments": [
                  {
                        "author_name": "Bob",
                        "fallback": "Bob did a thing",
                        "title": "Foo Created",
                        "id": 1,
                        "color": "ff0000",
                        "fields": [
                          {
                                "title": "ID",
                                "value": "65",
                                "short": true
                          }
                        ]
                  }
                ],
                "ts": "1522939727.000354"
          },
          {
                "type": "message",
                "user": "U1234",
                "text": "<@%s> hey bot",
                "ts": "1522939726.000713"
          }
        ],
        "has_more": false,
        "pin_count": 0
  }`, defaultNonBotUserID, defaultNonBotUserID, defaultNonBotUserID, defaultNonBotUserID, defaultBotID, defaultBotID, defaultBotID)
