// simpleserver project main.go
package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/sidmutha/simpleserver/myhttp"
)

func test() {

	//	var hmsg myhttp.Http_message = *myhttp.NewHttp_message()
	//	hmsg.SetStatus("HTTP/1.1 200 OK")
	//	hmsg.AddHeader("Date", "Mon, 27 Jul 2009 12:28:53 GMT")
	//	hmsg.AddHeader("Server", "Apache/2.2.14 (Win32)")
	//	hmsg.AddHeader("Last-Modified", "Wed, 22 Jul 2009 19:15:56 GMT")
	//	hmsg.AddHeader("Content-Type", "text/html")
	//	hmsg.SetBody("<html><body><h1>Hello</h1></body></html>")

	//	fmt.Println(hmsg)

	s := `GET /favicon.ico HTTP/1.1
Host: localhost:8188
Connection: keep-alive
User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36
Accept: image/webp,image/*,*/*;q=0.8
Referer: http://localhost:8188/p1
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: en-GB,en-US;q=0.8,en;q=0.6
Cookie: m=34e2:|77cb:t
If-Modified-Since: Wed, 22 Jul 2009 19:15:56 GMT

:t`
	hm := myhttp.ParseHttpMessage(s)
	v, _, _ := hm.GetStatus()
	fmt.Println(v)

}

func main() {

	port := ":8188"

	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Listening on port" + string(port))
	//	buf := make([]byte, 4096)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Accepted from " + conn.RemoteAddr().String())
		fmt.Println("-------------message------------")
		//conn.Read(buf)
		//fmt.Println(string(buf))
		//fmt.Println("***********/message*************")
		//conn.Write([]byte("<html><body><h1>Hello</h1></body></html>"))
		/*conn.Write([]byte(`HTTP/1.1 200 OK
		Date: Mon, 27 Jul 2009 12:28:53 GMT
		Server: Apache/2.2.14 (Win32)
		Last-Modified: Wed, 22 Jul 2009 19:15:56 GMT
		Content-Length: 52
		Content-Type: text/html
		Connection: Closed

		<html>
		<body>
		<h1>Hello, World!</h1>
		</body>
		</html>`))*/
		//conn.Write([]byte(hmsg.String()))
		//conn.Close()
		go myhttp.HandleConn(conn)
	}
}

func handleConnection(c net.Conn) {
	buf := make([]byte, 4096)

	for {
		n, err := c.Read(buf)
		if err != nil || n == 0 {
			c.Close()
			break
		}

		n, err = c.Write([]byte("<html><body><h1>Hello</h1></body></html>"))
		if err != nil {
			c.Close()
			break
		}
	}
	fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
}
