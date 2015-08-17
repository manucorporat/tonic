package simplejson

import (
	"bytes"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/manucorporat/tonic/common"
)

type conn struct {
	mutex  sync.Mutex
	socket *websocket.Conn
}

func newConn(socket *websocket.Conn) *conn {
	return &conn{socket: socket}
}

func (c *conn) Mutex() *sync.Mutex {
	return &c.mutex
}

func (c *conn) Send(msg common.Message) error {
	buf := new(bytes.Buffer)
	if err := encodeMsg(buf, msg); err != nil {
		return err
	}
	return c.socket.WriteMessage(websocket.TextMessage, buf.Bytes())
}

func (c *conn) Emit(eventName string, data interface{}) error {
	return c.Send(common.NewMsg(eventName, "", "", data))
}

func (c *conn) Recv() (common.Message, error) {
	for {
		msgType, reader, err := c.socket.NextReader()
		if err != nil {
			return nil, err
		}
		if msgType != websocket.TextMessage {
			continue
		}
		msg, err := decodeMsg(reader)
		if err != nil {
			c.socket.Close()
			return nil, err
		}
		return msg, nil
	}
}
