package messages

type Message struct {
	Room         string
	FromUserId   string
	FromUsername string
	ToUserId     string
	ToUsername   string
	Message      string
	Direct       bool
}
