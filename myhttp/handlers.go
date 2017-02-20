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

func GetExistFilePath(fpath string) (filepath string, err error) {
	f, e := os.Stat(fpath)
	filepath = fpath
	err = nil
	if e != nil {
		err = e
		return
	}
	if f.Mode().IsDir() {
		filepath += "/index.html"
		fmt.Println("folder requested, looking for index.html")
		return GetExistFilePath(filepath)
	}
	return
	//	switch mode := f.Mode() { // if?
	//		case mode.IsDir():
	//			filepath += "/index.html"
	//			return GetExistFilePath(filepath)
	////		case mode.IsRegular():
	////			filepath
	//	}
}

func HandleGET(m *Http_message) *Http_message {
	fmt.Println("GET: ")
	fmt.Println(m)
	fmt.Println("---")

	response := NewHttp_message()

	filepath, err := GetExistFilePath(rootdir + path)

	if err != nil {
		if os.IsNotExist(err) { // file doesn't exist, 404
			response.SetStatus("HTTP/1.1 404 Not Found Dude")
		} else { // some other error, 500
			response.SetStatus("HTTP/1.1 500 Internal Server Error")
		}
		// set custom 404 page?
		return response
	}

	data, err := ioutil.ReadFile(filepath)

	if err != nil { // error reading file
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
	response.SetBody(string(data))
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
		response_str := response.String()
		conn.Write([]byte(response_str))
		fmt.Println("Sending response: \n" + response_str + "---\n")
	}

	// keep-alive : time out before closing?
	conn.Close()
}
