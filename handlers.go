package slacktest

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	websocket "github.com/gorilla/websocket"
	slack "github.com/nlopes/slack"
)

func contextHandler(server *Server, next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ServerURLContextKey, server.GetAPIURL())
		ctx = context.WithValue(ctx, ServerWSContextKey, server.GetWSURL())
		ctx = context.WithValue(ctx, ServerBotNameContextKey, server.BotName)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body: %s", err.Error())
		return
	}
	log.Printf("Got a message posted: %s", string(data))
	values, vErr := url.ParseQuery(string(data))
	if vErr != nil {
		log.Printf("Unable to decode query params: %s", err.Error())
		return
	}

	ts := time.Now().Unix()
	resp := fmt.Sprintf(`{"channel":%s,"ts":%d, "text":"%s", "ok": true}`, values.Get("channel"), ts, values.Get("text"))
	m := slack.Message{}
	m.Type = "message"
	m.Channel = values.Get("channel")
	m.Timestamp = fmt.Sprintf("%d", ts)
	m.Text = values.Get("text")
	jsonMessage, jsonErr := json.Marshal(m)
	if jsonErr != nil {
		log.Printf("Unable to marshall message: %s", jsonErr.Error())
	} else {
		sendToChannel(string(jsonMessage))
	}
	_, _ = w.Write([]byte(resp))
}

func rtmStartHandler(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %s", err.Error())
		return
	}
	ts := slack.JSONTime(1402463766)
	wsurl := r.Context().Value(ServerWSContextKey).(string)
	rtmInfo := slack.Info{
		URL: wsurl,
		Team: &slack.Team{
			ID:     "T024BE7LD",
			Name:   "test slackbot team",
			Domain: "testdomain.com",
		},
		User: &slack.UserDetails{
			ID:             BotIDFromContext(r.Context()),
			Name:           BotNameFromContext(r.Context()),
			Created:        ts,
			ManualPresence: "true",
			Prefs:          slack.UserPrefs{},
		},
	}
	webResponse := slack.WebResponse{
		Ok: true,
	}
	fullresponse := fullInfoSlackResponse{
		rtmInfo,
		webResponse,
	}
	j, jErr := json.Marshal(fullresponse)
	if jErr != nil {
		msg := fmt.Sprintf("Unable to marshal response: %s", jErr.Error())
		log.Printf("Error: %s", msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	_, wErr := w.Write(j)
	if wErr != nil {
		log.Printf("Error writing response: %s", wErr.Error())
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Unable to upgrade to ws connection: %s", err.Error())
		return
	}
	defer func() { _ = c.Close() }()
	go handlePendingMessages(c)
	for {
		mt, messageBytes, err := c.ReadMessage()
		if err != nil {
			log.Printf("read error: %s", err.Error())
			continue
		}
		message := string(messageBytes)

		evt := &slack.Event{}
		if err := json.Unmarshal(messageBytes, evt); err != nil {
			log.Printf("Error unmarshalling message: %s", err.Error())
			log.Printf("failed message: %s", string(message))
			continue
		}
		if evt.Type == "ping" {
			p := &slack.Ping{}
			jErr := json.Unmarshal(messageBytes, p)
			if jErr != nil {
				log.Printf("Unable to decode ping event: %s", jErr.Error())
				continue
			}
			log.Print("responding to slack ping")
			pong := &slack.Pong{
				ReplyTo: p.ID,
				Type:    "pong",
			}
			j, _ := json.Marshal(pong)
			wErr := c.WriteMessage(mt, j)
			if wErr != nil {
				log.Printf("error writing pong back to socket: %s", wErr.Error())
				continue
			}
			continue
		} else {
			seenMessages = append(seenMessages, message)
			seenMessageChannel <- message
		}
	}
}
