package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	queue := make(chan chan string)

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

func client(q chan chan string, method string) {
	conn := make(chan string)
	q <- conn
	for i := 0; i < 10; i++ {
		//anObj.add(i)
		conn <- method + " /hello.html HTTP/1.1"
	}
	close(conn)
}

func server(q chan chan string, clients int) {
	for i := 0; i < clients; i++ {
		conn := <-q
		go handle(conn)
	}
}

func handle(conn chan string) {
	for req := range conn {
		fmt.Println("Got:", req)
	}
}
