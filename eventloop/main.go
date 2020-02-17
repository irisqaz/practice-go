package main

import (
	"fmt"
	"sync"
)

func main() {
	var wgSend sync.WaitGroup
	var wgReceive sync.WaitGroup
	queue := make(chan func())

	go func() {
		runFuncs(queue, func() {
			fmt.Println("Hello World")
		})
		wgSend.Done()
	}()

	go func() {
		runFuncs(queue, func() {
			fmt.Println("Bye World")
		})
		wgSend.Done()
	}()

	go func() {
		receiveFuncs(queue)
		wgReceive.Done()
	}()

	wgSend.Add(2)
	wgReceive.Add(1)

	go func() {
		wgSend.Wait()
		close(queue)
	}()

	wgReceive.Wait()

}

func runFuncs(q chan func(), f func()) {
	for i := 0; i < 10; i++ {
		//anObj.add(i)
		q <- f
	}
}

func receiveFuncs(q chan func()) {
	// for  {
	// 	value := <- q
	// 	anObj.add(value)
	// }
	for f := range q {
		f()
	}
}
