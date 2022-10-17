package main

import (
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "TCP server: ", 0)

	l.Println("Hello world!")
}
