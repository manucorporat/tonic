package common

import "net/http"

type ConnType int

const (
	FullDuplex     ConnType = 0
	ServerToClient ConnType = 1
	ClientToServer ConnType = 2
)

type Protocol interface {
	Upgrade(http.ResponseWriter, *http.Request) Connection
}

type Connection interface {
	Send(Message) error
	Emit(string, interface{}) error
	Recv() (Message, error)
}

type Message interface {
	Namespace() string
	Name() string
	Id() string
	Data() interface{}
}
