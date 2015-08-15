package sseio

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Upgrader struct {
	Upgrader websocket.Upgrader
}

func Default() *Upgrader {
	return NewWithUpgrader(websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})
}

func NewWithUpgrader(upgrader websocket.Upgrader) *Upgrader {
	return &Upgrader{
		Upgrader: upgrader,
	}
}

func (u *Upgrader) Upgrade(w http.ResponseWriter, req *http.Request) (*conn, error) {
	conn, err := u.Upgrader.Upgrade(w, req, nil)
	if err != nil {
		return nil, err
	}
	return newConn(conn), nil
}
