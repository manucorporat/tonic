package common

func Emit(conn Connection, eventName string, data interface{}) error {
	return conn.Send(NewMsg(eventName, "", "", data))
}
