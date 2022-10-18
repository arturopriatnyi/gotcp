package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

var l = log.New(os.Stdout, "TCP server: ", 0)

func main() {
	// getting server's port
	if len(os.Args) == 1 {
		l.Fatalln("missing port")
	}
	port := ":" + os.Args[1]

	// starting the server
	s, err := net.Listen("tcp", port)
	if err != nil {
		l.Fatalf("server starting failed: %v\n", err)
	}
	l.Println("server started")

	// deferring server closing
	defer func(s net.Listener) {
		if err = s.Close(); err != nil {
			l.Fatalf("server closing failed: %v\n", err)
		}
		l.Println("server closed")
	}(s)

	for {
		// accepting a client
		c, err := s.Accept()
		if err != nil {
			l.Printf("client acceptance failed: %v\n", err)
		}

		// handling the client concurrently
		go handle(c)
	}
}

func handle(c net.Conn) {
	// request
	req, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		l.Printf("request reading failed: %v\n", err)
	}
	l.Printf("request: %s", req)

	// response
	if _, err = c.Write([]byte(req)); err != nil {
		l.Printf("response sending failed: %v\n", err)
	}
	l.Printf("response: %s", req)

	// closing client connection
	if err = c.Close(); err != nil {
		l.Printf("client connection closing failed: %v\n", err)
	}
}
