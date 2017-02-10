package myhttp

import (
	"strings"
)

func (m *Http_message) GetStatus() (method string, path string, proto string) {
	s := strings.SplitN(m.status, " ", 3)
	method = s[0]
	path = s[1]
	proto = s[2]
	return
}
