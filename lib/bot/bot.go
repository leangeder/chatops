package bot

import (
  "fmt"
  "log"
  "strings"
  "github.com/leangeder/chatops/lib/message"
  "sync"
)

type Bot struct {
  name            string
  receivedMessage chan messages.Message
  sendMessage     chan messages.Message
}

var processMother sync.Once

func New(config map[string]interface{}) *Bot {
  bot := &Bot{
    name:               config["username"].(string),
    receivedMessage:     make(chan messages.Message),
    sendMessage:         make(chan messages.Message),
  }
  log.Println("bot: ", bot.name)

  return bot
}

func (bot *Bot) Process() {
  processMother.Do(func() {
    log.Println("bot: starting main loop")
    for rmessage := range bot.receivedMessage {
      if strings.HasPrefix(rmessage.Message, bot.name + " help") {
        log.Println("bot: starting main loop")
        go func(b Bot, msg messages.Message) {
          helpMsg := fmt.Sprintln("Available Commands:")
          // for _, rule := range s.rules {
          //   helpMsg = fmt.Sprintln(helpMsg, rule.HelpMessage(b, bot.Room))
          // }
          bot.sendMessage <- messages.Message{
            Room:                msg.Room,
            ReceiverUserID:      msg.SenderUserID,
            ReceiverUserName:    msg.SenderUserName,
            Message:             helpMsg,
          }
        }(*bot, rmessage)
        continue
      }

      go func(bot Bot, msg messages.Message) {
        defer func() {
          if r := recover(); r != nil {
            log.Printf("panic recovered when parsing message: %#v. Panic: %v", msg, r)
          }
        }()
      }(*bot, rmessage)
    }
  })
}
