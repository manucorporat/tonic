package common

import (
	"encoding/json"
	"fmt"
)

type EmptyMsg struct{}

var _ Message = EmptyMsg{}

func (EmptyMsg) Namespace() string {
	return ""
}

func (EmptyMsg) Name() string {
	panic("Event name must be implemented")
}

func (EmptyMsg) Id() string {
	return ""
}

func (EmptyMsg) Data() []byte {
	return nil
}

type Msg struct {
	name      string
	data      []byte
	id        string
	namespace string
}

var _ Message = Msg{}

func NewMsg(name, id, namespace string, data []byte) Msg {
	return Msg{
		name:      name,
		id:        id,
		namespace: namespace,
		data:      data,
	}
}

func (m Msg) Namespace() string {
	return m.namespace
}

func (m Msg) Name() string {
	return m.name
}

func (m Msg) Id() string {
	return m.id
}

func (m Msg) Data() []byte {
	return m.data
}

func (m Msg) String() string {
	return fmt.Sprintf(
		`Message(%s):
	- Id: %s
	- Namespace: %s
	- Data: %v`, m.name, m.id, m.namespace, m.data)
}

func BindJSON(msg Message, obj interface{}) error {
	return json.Unmarshal(msg.Data(), obj)
}
