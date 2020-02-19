package main

import (
	"fmt"
	"sync"
)

type message string
type connection chan message

func main() {
	var wg sync.WaitGroup
	queue := make(chan connection)

	go func() {
		client(queue, "GET")
		wg.Done()
	}()

	go func() {
		client(queue, "POST")
		wg.Done()
	}()

	server(queue, 2)

	wg.Add(2)
	wg.Wait()
}

func client(q chan connection, method string) {
	conn := make(connection)
	q <- conn
	for i := 0; i < 10; i++ {
		//anObj.add(i)
		msg := message(method + " /hello.html HTTP/1.1")
		conn <- msg
	}
	close(conn)
}

func server(q chan connection, clients int) {
	for i := 0; i < clients; i++ {
		conn := <-q
		go handle(conn)
	}
}

func handle(conn connection) {
	for msg := range conn {
		fmt.Println("Got:", msg)
	}
}
