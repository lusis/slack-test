# slacktest
This is a very basic golang library for testing your slack RTM chatbots

Depending on your mechanism for building slackbots in go and how broken out your message parsing logic is, you can normally test that part pretty cleanly.
However testing your bots RESPONSES are a bit harder.

This library attempts to make that a tad bit easier but in a slightly opinionated way.

The current most popular slack library for golang is [nlopes/slack](https://github.com/nlopes/slack). Conviently the author has made overriding the slack API endpoint a feature. This allows us to use our fake slack server to inspect the chat process.

## Limitations
Right now the test server is VERY limited. It currently handles the following two API endpoints

- `rtm.start`
- `chat.postMessage`

This is enough to do the initial testing I wanted to be able to accomplish.

## Example usage
You can see an example in the `examples` directory of how to you might test it

If you just want to play around:

```go
package main

import (
    "log"
    slacktest "github.com/lusis/slack-test"
    slackbot "github.com/lusis/go-slackbot"
    slack "github.com/nlopes/slack"
)

func globalMessageHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "I see your message", slackbot.WithoutTyping)
}

func main() {
    // Setup our test server
    s := slacktest.NewTestServer()
    // Set a custom name for our bot
    s.SetBotName("MyBotName")
    // ensure that anything using the slack api library uses our custom server
    slack.SLACK_API = "http://" + s.ServerAddr + "/"
    // start the test server
    go s.Start()
    // create a new slackbot. Token is irrelevant here
    bot := slackbot.New("ABCEDFG")
    // add a message handler
    bot.Hear("this is a channel message").MessageHandler(globalMessageHandler)
    // start the bot
    go bot.Run()
    //send a message to a channel
    s.SendMessageToChannel("#random", "this is a channel message")
    for m := range s.SeenFeed {
        log.Printf("saw message in slack: %s", m)
    }
}
```
Output:
```
# go run main.go
2017/09/26 10:53:49 {"type":"message","channel":"#random","text":"this is a channel message","ts":"1506437629","pinned_to":null}
2017/09/26 10:53:49 {"id":1,"channel":"#random","text":"I see your message","type":"message"}
#
```
You can see that our bot handled the message correctly.

## Usage in tests
Currently it's not as ergonomic as I'd like. So much depends on how modular your bot code is in being able to run the same message handling code against a test instance. In the `examples` directory there are a couple of test cases.

Additionally, you definitely need to call `time.Sleep` for now in your test code to give time for the messages to work through the various channels and populate. I'd like to add a safer subscription mechanism in the future with a proper timeout mechanism.
