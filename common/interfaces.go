package common

import (
	"net/http"
	"sync"
)

type Protocol interface {
	Upgrade(http.ResponseWriter, *http.Request) Connection
}

type Connection interface {
	Mutex() *sync.Mutex
	Send(Message) error
	Recv() (Message, error)
}

type Message interface {
	Namespace() string
	Name() string
	Id() string
	Data() interface{}
}
