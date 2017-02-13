package myhttp

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

var rootdir, path string

func FileExists(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}

func HandleGET(m *Http_message) *Http_message {
	fmt.Println("Handle GET received: ")
	fmt.Println(m)
	fmt.Println("------------------")

	response := NewHttp_message()

	filepath := rootdir + path
	if !FileExists(filepath) {
		response.SetStatus("HTTP/1.1 404 Not Found Dude")
		// set custom 404 page?
		return response
	}

	dat, err := ioutil.ReadFile(filepath)

	if err != nil {
		response.SetStatus("HTTP/1.1 500 Internal Server Error")
		// set custom 500 page?
		return response
	}
	response.SetStatus("HTTP/1.1 200 OK")
	//response.AddHeader("Date", "Mon, 27 Jul 2009 12:28:53 GMT")
	//response.AddHeader("Server", "Apache/2.2.14 (Win32)")
	//response.AddHeader("Last-Modified", "Wed, 22 Jul 2009 19:15:56 GMT")
	response.AddHeader("Content-Type", "text/html")
	//response.SetBody("<html><body><h1>Hi</h1></body></html>")
	response.SetBody(string(dat))
	return response
}

func HandleConn(conn net.Conn, rdir string) {
	rootdir = rdir
	buf := make([]byte, 4096)
	// only this handler should send/recv over the conn
	conn.Read(buf)
	fmt.Println(string(buf))
	http_message := ParseHttpMessage(string(buf))
	verb, p, _ := http_message.GetStatus()

	path = p

	// channel might help when handling keep-alive connections, logging, etc
	//c := make(chan Http_message)

	if verb == "GET" {
		response := HandleGET(http_message)
		conn.Write([]byte(response.String()))
	}

	// keep-alive : time out before closing?
	conn.Close()
}
