package plugins

import (
  "log"
  "github.com/leangeder/chatops/lib/message"
)

type Plugin interface {
  ReceivedMessage() chan messages.Message
  SendMessage()     chan messages.Message
  Error()           error
}

var listPlugins []func(func(string) string) (Plugin, bool)

func Load() Plugin {

  return CLI()
}
