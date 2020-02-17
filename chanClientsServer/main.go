package main

import (
	"sync"

	"github.com/irisqaz/practice-go/net"
)

var count int

func main() {
	var wg sync.WaitGroup
	queue := make(chan chan string)
	go func() {
		net.Client(queue, "GET")
		wg.Done()
	}()

	go func() {
		net.Client(queue, "POST")
		wg.Done()
	}()

	net.Server(queue, 2)

	wg.Add(2)
	wg.Wait()
}
