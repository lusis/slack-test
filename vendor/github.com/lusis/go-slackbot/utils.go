package slackbot

import (
	"regexp"

	"github.com/nlopes/slack"
)

// StripDirectMention removes a leading mention (aka direct mention) from a message string
func StripDirectMention(text string) string {
	r, rErr := regexp.Compile(`(^<@[a-zA-Z0-9]+>[\:]*[\s]*)?(.*)`)
	if rErr != nil {
		return ""
	}
	return r.FindStringSubmatch(text)[2]
}

// IsDirectMessage returns true if this message is in a direct message conversation
func IsDirectMessage(evt *slack.MessageEvent) bool {
	r, rErr := regexp.Compile("^D.*")
	if rErr != nil {
		return false
	}
	return r.MatchString(evt.Channel)
}

// IsDirectMention returns true is message is a Direct Mention that mentions a specific user. A
// direct mention is a mention at the very beginning of the message
func IsDirectMention(evt *slack.MessageEvent, userID string) bool {
	r, rErr := regexp.Compile("^<@" + userID + ">.*")
	if rErr != nil {
		return false
	}
	return r.MatchString(evt.Text)
}

// IsMentioned returns true if this message contains a mention of a specific user
func IsMentioned(evt *slack.MessageEvent, userID string) bool {
	userIDs := WhoMentioned(evt)
	for _, u := range userIDs {
		if u == userID {
			return true
		}
	}
	return false
}

// IsMention returns true the message contains a mention
func IsMention(evt *slack.MessageEvent) bool {
	r, rErr := regexp.Compile(`<@(U[a-zA-Z0-9]+)>`)
	if rErr != nil {
		return false
	}
	results := r.FindAllStringSubmatch(evt.Text, -1)
	return len(results) > 0
}

// WhoMentioned returns a list of userIDs mentioned in the message
func WhoMentioned(evt *slack.MessageEvent) []string {
	r, rErr := regexp.Compile(`<@(U[a-zA-Z0-9]+)>`)
	if rErr != nil {
		return []string{}
	}
	results := r.FindAllStringSubmatch(evt.Text, -1)
	matches := make([]string, len(results))
	for i, r := range results { // nolint: gosimple
		matches[i] = r[1]
	}
	return matches
}

func namedRegexpParse(message string, exp *regexp.Regexp) (bool, map[string]string) {
	md := make(map[string]string)
	allMatches := exp.FindStringSubmatch(message)
	if len(allMatches) == 0 {
		return false, md
	}
	keys := exp.SubexpNames()
	if len(keys) != 0 {
		for i, name := range keys {
			if i != 0 {
				md[name] = allMatches[i]
			}
		}
	}
	return true, md
}
