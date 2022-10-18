package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	l := log.New(os.Stdout, "TCP client: ", 0)

	// getting server's IP address
	if len(os.Args) == 1 {
		l.Fatalln("missing server IP address")
	}
	addr := os.Args[1]

	// connecting to a server
	c, err := net.Dial("tcp", addr)
	if err != nil {
		l.Fatalf("server connection failed: %s\n", addr)
	}
	l.Println("client connected")

	// deferring connection closing
	defer func(c net.Conn) {
		if err = c.Close(); err != nil {
			l.Fatalf("client connection closing failed: %v\n", err)
		}
		l.Println("client connection closed")
	}(c)

	// request
	req := []byte("ping\n")
	if _, err = c.Write(req); err != nil {
		l.Printf("request sending failed: %v\n")
	}
	l.Printf("request: %s", req)

	// response
	res, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		l.Printf("response reading failed: %v\n", err)
	}
	l.Printf("response: %s", res)
}
