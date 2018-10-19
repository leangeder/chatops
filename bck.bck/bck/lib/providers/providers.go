package provider

type Provider interface {
	IncomingChannel() chan messages.Message
	OutgoingChannel() chan messages.Message
	Error() error
}

func (c *providerCLI) IncomingChannel() chan messages.Message {
	return c.in
}

func (c *providerCLI) OutgoingChannel() chan messages.Message {
	return c.out
}

func (c *providerCLI) Error() error {
	return nil
}
