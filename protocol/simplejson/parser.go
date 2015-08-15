package simplejson

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"

	"github.com/manucorporat/tonic/common"
)

var escaper = strings.NewReplacer(":", "%3A")
var unescaper = strings.NewReplacer("%3A", ":")

func encodeMsg(w *bytes.Buffer, msg common.Message) error {
	w.WriteString(escaper.Replace(msg.Name()))
	w.WriteRune(':')
	w.WriteString(escaper.Replace(msg.Id()))
	w.WriteRune(':')
	w.WriteString(escaper.Replace(msg.Namespace()))
	w.WriteRune(':')

	data := msg.Data()
	switch data.(type) {
	case string:
		w.WriteString(data.(string))
	case []byte:
		w.Write(data.([]byte))
	default:
		return json.NewEncoder(w).Encode(msg.Data())

	}
	return nil
}

func decodeMsg(reader io.Reader) (common.Message, error) {
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
