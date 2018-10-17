package messages

type Message struct {
  Room             string
  SenderUserID     string
  SenderUserName   string
  ReceiverUserID   string
  ReceiverUserName string
  Message          string
  IsPrivate        bool
}
