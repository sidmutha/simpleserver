package myhttp

import (
	"fmt"
	"net"
)

func HandleGET(m *Http_message) *Http_message {
	fmt.Println("Handle GET received: ")
	fmt.Println(m)
	fmt.Println("------------------")

	response := NewHttp_message()
	response.SetStatus("HTTP/1.1 200 OK")
	response.AddHeader("Date", "Mon, 27 Jul 2009 12:28:53 GMT")
	response.AddHeader("Server", "Apache/2.2.14 (Win32)")
	response.AddHeader("Last-Modified", "Wed, 22 Jul 2009 19:15:56 GMT")
	response.AddHeader("Content-Type", "text/html")
	response.SetBody("<html><body><h1>Hi</h1></body></html>")
	return response
}

func HandleConn(conn net.Conn) {
	buf := make([]byte, 4096)
	conn.Read(buf)
	fmt.Println(string(buf))
	http_message := ParseHttpMessage(string(buf))
	verb, _, _ := http_message.GetStatus()
	// only this handler should send/recv over the conn

	// channel might help when handling keep-alive connections, logging, etc

	//c := make(chan Http_message)

	if verb == "GET" {
		response := HandleGET(http_message)
		conn.Write([]byte(response.String()))
	}

	// keep-alive : time out before closing?
	conn.Close()
}
