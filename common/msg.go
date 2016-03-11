package common

import (
	"encoding/json"
	"errors"
	"fmt"
)

type EmptyMsg struct{}

var _ Message = EmptyMsg{}

func (EmptyMsg) Namespace() string {
	return ""
}

func (EmptyMsg) Name() string {
	return ""
}

func (EmptyMsg) Id() string {
	return ""
}

func (EmptyMsg) Data() interface{} {
	return nil
}

type Msg struct {
	name      string
	data      interface{}
	id        string
	namespace string
}

var _ Message = Msg{}

func NewMsg(name, id, namespace string, data interface{}) Msg {
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

func (m Msg) Data() interface{} {
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
	data, ok := msg.Data().([]byte)
	if !ok {
		return errors.New("body can not be binded to json")
	}
	return json.Unmarshal(data, obj)
}
