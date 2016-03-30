package common

func BasicMsg(eventName string, data []byte) Msg {
	return NewMsg(eventName, "", "", data)
}

func Emit(conn Connection, eventName string, data []byte) error {
	return conn.Send(BasicMsg(eventName, data))
}
