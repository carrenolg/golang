package models

type Message struct {
	Name string
	Body string
	Time int64
}

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

type IncomingMessage struct {
	Cmd *Command
	Msg *Message
}

type Command struct {
	Type int
	Text string
}
