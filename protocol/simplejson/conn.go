package simplejson

import (
	"bufio"
	"io"
	"io/ioutil"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/manucorporat/tonic/common"
)

var escaper = strings.NewReplacer(":", "%3A")
var unescaper = strings.NewReplacer("%3A", ":")

type conn struct {
	socket *websocket.Conn
}

var _ common.Connection = &conn{}

func newConn(socket *websocket.Conn) *conn {
	return &conn{socket: socket}
}

func (c *conn) Send(msg common.Message) error {
	buf := c.encode(msg)
	return c.socket.WriteMessage(websocket.BinaryMessage, buf)
}

func (c *conn) Emit(eventName string, data []byte) error {
	return c.Send(common.NewMsg(eventName, "", "", data))
}

func (c *conn) Recv() (common.Message, error) {
	for {
		msgType, reader, err := c.socket.NextReader()
		if err != nil {
			return nil, err
		}
		if msgType != websocket.BinaryMessage {
			continue
		}
		msg, err := c.decode(reader)
		if err != nil {
			c.socket.Close()
			return nil, err
		}
		return msg, nil
	}
}

func (c *conn) Close() error {
	return c.socket.Close()
}

func (c *conn) encode(msg common.Message) []byte {
	name := escaper.Replace(msg.Name())
	id := escaper.Replace(msg.Id())
	namespace := escaper.Replace(msg.Namespace())
	data := msg.Data()

	totalsize := len(name) + len(id) + len(namespace) + len(data) + 3
	buf := make([]byte, totalsize)

	// event name
	n := copy(buf, name)
	buf[n] = ':'
	n++

	// id
	n += copy(buf[n:], id)
	buf[n] = ':'
	n++

	// namespace
	n += copy(buf[n:], namespace)
	buf[n] = ':'
	n++

	// data
	copy(buf[n:], data)
	return buf
}

func (c *conn) decode(reader io.Reader) (common.Message, error) {
	buf := bufio.NewReader(reader)
	eventName, err := buf.ReadString(':')
	if err != nil {
		return nil, err
	}
	id, err := buf.ReadString(':')
	if err != nil {
		return nil, err
	}
	namespace, err := buf.ReadString(':')
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(buf)
	if err != nil {
		return nil, err
	}

	return common.NewMsg(
		unescaper.Replace(eventName[:len(eventName)-1]),
		unescaper.Replace(id[:len(id)-1]),
		unescaper.Replace(namespace[:len(namespace)-1]),
		content), nil
}
