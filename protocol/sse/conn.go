package sseio

import (
	"bytes"

	"github.com/gorilla/websocket"
	"github.com/manucorporat/sse"
	"github.com/manucorporat/tonic/common"
)

type conn struct {
	socket  *websocket.Conn
	pending *common.Queue
}

func newConn(socket *websocket.Conn) *conn {
	return &conn{
		socket:  socket,
		pending: common.NewQueue(),
	}
}

func (c *conn) Type() common.ConnType {
	return common.FullDuplex
}

func (c *conn) Send(msg common.Message) error {
	buf := new(bytes.Buffer)
	err := sse.Encode(buf, sse.Event{
		Event: msg.Name(),
		Id:    msg.Id(),
		Data:  msg.Data(),
	})
	if err != nil {
		return err
	}
	return c.socket.WriteMessage(websocket.TextMessage, buf.Bytes())
}

func (c *conn) Recv() (common.Message, error) {
	for {
		msg := c.pending.Dequeue()
		if msg != nil {
			return msg.(common.Msg), nil
		}

		err := c.nextpackage()
		if err != nil {
			return nil, err
		}
	}
}

func (c *conn) nextpackage() error {
	msgType, reader, err := c.socket.NextReader()
	if err != nil {
		return err
	}
	if msgType != websocket.TextMessage {
		return nil
	}
	events, err := sse.Decode(reader)
	if err != nil {
		return err
	}
	for _, event := range events {
		c.pending.Enqueue(newMsg(event))
	}
	return nil
}

func newMsg(event sse.Event) common.Msg {
	return common.NewMsg(event.Event, event.Id, "", event.Data)
}
