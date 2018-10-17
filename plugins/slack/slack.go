package plugins

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/leangeder/chatops/lib/configuration"
	"golang.org/x/net/websocket"
)

const (
	urlSlackAPI = "https://slack.com/api/"
)

type slack struct {
	token    string
	wsURL    string
	selfID   string
	wsConnMu sync.Mutex
	wsConn   *websocket.Conn

	err error

	mu        sync.Mutex
	usernames map[string]string
}

func (p *slack) connect() {
	log.Println("slack: connecting to HTTP API handshake interface")
	resp, err := http.Get(fmt.Sprint(urlSlackAPI, "rtm.start?no_unreads&simple_latest&token=", p.token))
	if err != nil {
		p.err = err
		return
	}
	defer resp.Body.Close()
	var data struct {
		OK   interface{} `json:"ok"`
		URL  string      `json:"url"`
		Self struct {
			ID string `json:"id"`
		} `json:"self"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		p.err = err
		return
	}

	switch v := data.OK.(type) {
	case bool:
		if !v {
			p.err = err
			return
		}
	default:
		p.err = err
		return
	}
	p.wsURL = data.URL
	p.selfID = data.Self.ID

	// dial part
	log.Println("slack: dialing to HTTP WS rtm interface")
	if p.wsURL == "" {
		p.err = fmt.Errorf("could not connnect to Slack HTTP WS rtm. please, check your connection and your token (%s). error: %v", p.token, p.err)
		return
	}
	ws, err := websocket.Dial(p.wsURL, "", urlSlackAPI)
	if err != nil {
		p.err = err
		return
	}
	p.wsConnMu.Lock()
	p.wsConn = ws
	p.wsConnMu.Unlock()
}

func (p *slack) listen() {
	log.Println("slack: started message intake loop")
	for {
		var data struct {
			Type    string `json:"type"`
			Channel string `json:"channel"`
			UserID  string `json:"user"`
			Text    string `json:"text"`
		}

		p.wsConnMu.Lock()
		wsConn := p.wsConn
		p.wsConnMu.Unlock()

		if err := json.NewDecoder(wsConn).Decode(&data); err != nil {
			continue
		}

		if data.Type != "message" {
			continue
		}

		// // // msg := messages.Message{
		// // // 	Room:         data.Channel,
		// // // 	FromUserID:   data.UserID,
		// // // 	FromUserName: p.getUserName(data.UserID),
		// // // 	Message:      data.Text,
		// // // 	Direct:       strings.HasPrefix(data.Channel, "D"),
		// // // }
		// // msg struct {
		// // 	Room         data.Channel
		// // 	FromUserID   data.UserID
		// // 	FromUserName p.getUserName(data.UserID)
		// // 	Message      data.Text
		// // 	Direct       strings.HasPrefix(data.Channel, "D")
		// // }
		// // p.in <- msg
	}
}

func (p *slack) reconnect() {
	for {
		time.Sleep(1 * time.Second)

		p.wsConnMu.Lock()
		wsConn := p.wsConn
		p.wsConnMu.Unlock()

		if wsConn == nil {
			log.Println("slack: cannot reconnect")
			break
		}

		if _, err := wsConn.Write([]byte(`{"type":"hello"}`)); err != nil {
			log.Printf("slack: reconnecting (%v)", err)
			p.connect()
		}
	}
}

type Configuration struct {
	slack []struct {
		Token string `yaml:"token"`
		Url   string `yaml:"url"`
	} `yaml:"slack"`
}

type PluginConfiguration struct {
	configuration *configuration.Configuration
	// name   string "slack"
	// pseudo string ""
	// token  string ""
	// url    string ""
	// // slack  []struct {
	// // 	Token string `yaml:"token"`
	// // 	Url   string `yaml:"url"`
	// // } `yaml:"slack"`
}

func New(config *configuration.Configuration) (*PluginConfiguration, error) {

	if config == nil {
		return nil, errors.New("Invalid configuration")
	}

	return &PluginConfiguration{configuration: config}, nil
	// return &PluginConfiguration{pseudo: config.Pseudo, token: config.Slack.token, url: config.Slack.url}, nil

	// p := &slack{
	// 	token: sxoxb-451023708149-453324399522-VceD40ITKKuSfN1mWQk4vip1",
	// 	// token: config["username"].(string),
	// 	// token:     config["chats"].["slack"].["token"].(string),
	// 	usernames: make(map[string]string),
	// }

	// p.connect()
	// if p.err == nil {
	// 	go p.listen()
	// }
	// p.reconnect()
}
