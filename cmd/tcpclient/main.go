package main

import (
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "TCP client: ", 0)

	if len(os.Args) == 1 {
		l.Fatalln("port is not provided")
	}

	port := os.Args[1]
	l.Println("port:", port)
}
