package message

type Message struct {
	Room         string
	FromUserID   string
	FromUserName string
	ToUserID     string
	ToUserName   string
	Message      string
	Direct       bool
}
