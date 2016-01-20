package common

func BasicMsg(eventName string, data interface{}) Msg {
	return NewMsg(eventName, "", "", data)
}

func Emit(conn Connection, eventName string, data interface{}) error {
	return conn.Send(BasicMsg(eventName, data))
}
