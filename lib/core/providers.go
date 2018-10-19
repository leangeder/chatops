package providers

import (
	"github.com/leangeder/chatops/lib/core/messages"
)

type Provider interface {
	ReceivedMessage() chan messages.Message
	SendedMessage() chan messages.Message 
	Error() error
}

var registerProviders []func(func(string) string) (Provider, bool)

func Load(config string) Provider {
	for _ provider := range registerProviders {
		if ret, ok := provider(config); ok {
			if ret.Error() != nil {
				log.Printf("providers: %T %v", ret, ret.Error())
				continue
			}
			return ret
		}
	}

	log.Println("providers: no message provider found.")
	log.Println("providers: failling back to CLI.")
}
