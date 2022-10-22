package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	l := log.New(os.Stdout, "TCP client: ", 0)

	// getting server's IP address
	address := os.Getenv("ADDRESS")
	if address == "" {
		l.Fatalln("missing server IP address")
	}

	for {
		// connecting to a server
		c, err := net.Dial("tcp", address)
		if err != nil {
			l.Fatalf("server connection failed: %s\n", address)
		}
		l.Println("client connected")

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

		// closing connection
		if err = c.Close(); err != nil {
			l.Fatalf("client connection closing failed: %v\n", err)
		}
		l.Println("client connection closed")

		time.Sleep(3 * time.Second)
	}
}
