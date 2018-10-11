package slack


import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
  "time"

	"golang.org/x/net/websocket"
)
type providerSlack struct {
	token		 	   string
	wsURL   		 string
	selfID  		 string
	wsConnMu		 sync.Mutex
	wsConn  		 *websocket.Conn

	in  chan		 messages.Message
	out chan		 messages.Message
	err error

	mu       		 sync.Mutex
	usernames		 map[string]string
}

func init() {
	token :=
	if token == "" {
		log.Println("providers: skipping Slack. if you want Slack enabled, please set a valid value for the environment variables", slackEnvVarName)
		return nil, false
	}
	return Slack(token), true
}


//
// import "fmt"
//
// func loadConfig(config map[string]interface{}) err {
//
// 	token, err := config["chats"]["slack"]["token"]
// 	if err != nil {
// 		panic(fmt.Println("Unable to load token for slack plugin: %s", err))
// 	}
//
// 	url, err := config["chats"]["slack"]["url"]
// 	if err != nil {
// 		panic(fmt.Println("Unable to found url for slack plugin: %s", err))
// 	}
//
// 	return c
// }
//
// func Run() {
// 	slackClient = slack.New(config["chats"]["slack"]["token"])
// 	rtm := slackClient.NewRTM()
// 	go rtm.ManageConnection()
//
// 	for msg := range rtm.IncomingEvents {
// 		switch ev := msg.Data.(type) {
// 			case *slack.MessageEvent:
// 				if len(ev.BotID) == 0 {
// 					go handleMessage(ev)
// 				}
// 			}
// 		}
// 	}
// }
//
// func handleMessage(ev *slack.MessageEvent) {
// 	result, err := wit
// }
//
// func replyToUser(ev *slack.MessageEvent, message string) {
// 	slackClient.PostMessage(ev.User, message, slack.PostMessageParameters{
// 		AsUser: true,
// 	})
// }
