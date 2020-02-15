package net

import (
	"fmt"
)

func Client(q chan chan string, method string) {
	conn := make(chan string)
	q <- conn
	for i := 0; i < 1000; i++ {
		conn <- method + " /hello.html HTTP/1.1"
	}
	close(conn)
}

func Server(queue chan chan string, clients int) {
	for i := 0; i < clients; i++ {
		conn := <-queue
		go handle(conn)
	}
}

func handle(conn chan string) {
	for req := range conn {
		fmt.Println("got:", req)
	}
}
