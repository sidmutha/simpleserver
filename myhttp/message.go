package myhttp

import (
	"strings"
)

type Http_message struct {
	status  string
	headers map[string]string
	body    string
}

func (m *Http_message) SetStatus(status string) {
	m.status = status
}

func (m *Http_message) SetBody(body string) {
	m.body = body
}

func (m *Http_message) AddHeader(key string, val string) {
	// add to map
	m.headers[key] = val
}

func NewHttp_message() *Http_message {
	return &Http_message{headers: make(map[string]string)}
}

func (m Http_message) String() string {
	var msg_string string
	msg_string += m.status
	msg_string += "\n"
	for k, v := range m.headers {
		msg_string += k + " : " + v + "\n"
	}
	msg_string += "\n"
	msg_string += m.body
	return msg_string
}

func ParseHttpMessage(message string) *Http_message {
	hmsg := NewHttp_message()
	msg_arr := strings.Split(message, "\n")
	hmsg.SetStatus(msg_arr[0])
	var i int
	i = 0
	for _, m := range msg_arr[1:] {
		if m == "" {
			break
		}
		kv := strings.SplitN(m, ":", 2)
		key := strings.TrimSpace(kv[0])
		val := strings.TrimSpace(kv[1])
		hmsg.AddHeader(key, val)
	}
	hmsg.SetBody(strings.Join(msg_arr[i:], "\n"))
	return hmsg
}
