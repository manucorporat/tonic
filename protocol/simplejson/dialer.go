package simplejson

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func Dial(urlStr string, requestHeader http.Header) (*conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(urlStr, requestHeader)
	if err != nil {
		return nil, err
	}
	return newConn(conn), nil
}
