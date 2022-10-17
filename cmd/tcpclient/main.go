package main

import (
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "TCP client: ", 0)

	l.Println("Hello world!")
}
