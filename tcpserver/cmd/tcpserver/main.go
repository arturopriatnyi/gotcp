package main

import (
	"bufio"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
)

const address = ":8080"

var l = log.New(os.Stdout, "TCP server: ", 0)

func main() {
	// starting the server
	s, err := net.Listen("tcp", address)
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

	// graceful shutdown setup
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	var wg sync.WaitGroup

	go func() {
		for {
			// accepting a client
			c, err := s.Accept()
			if errors.Is(err, net.ErrClosed) {
				break
			}
			if err != nil {
				l.Printf("client acceptance failed: %v\n", err)
			}

			// handling the client concurrently
			go func() {
				wg.Add(1) // adding new client to the counter
				handle(c)
				wg.Done() // releasing a client
			}()
		}
	}()

	<-quit    // waiting for interrupt
	wg.Wait() // waiting until all accepted clients are handled
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
