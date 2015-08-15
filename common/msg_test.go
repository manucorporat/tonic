package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMsg(t *testing.T) {
	msg := Msg{
		MsgName:      "event",
		MsgData:      12,
		MsgId:        "1234",
		MsgNamespace: "/path",
	}
	assert.Equal(t, msg.Name(), "event")
	assert.Equal(t, msg.Data(), 12)
	assert.Equal(t, msg.Id(), "1234")
	assert.Equal(t, msg.MsgNamespace(), "/path")
}
