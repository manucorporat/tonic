package common

import "net/http"

type Protocol interface {
	Upgrade(http.ResponseWriter, *http.Request) Connection
}

type Connection interface {
	Send(Message) error
	Recv() (Message, error)
}

type Message interface {
	Namespace() string
	Name() string
	Id() string
	Data() []byte
}
