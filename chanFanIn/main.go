package main

import (
	"fmt"
	"sync"
)

type sharedObj struct {
	records []int
}

func (obj *sharedObj) add(n int) {
	obj.records = append(obj.records, n)

}

func (obj *sharedObj) getRecords() []int {
	return obj.records
}

var anObj sharedObj

func main() {
	var wg sync.WaitGroup
	var wgReceive sync.WaitGroup
	queue := make(chan int)

	go func() {
		runWrites(0, 10, queue)
		wg.Done()
	}()

	go func() {
		runWrites(10, 20, queue)
		wg.Done()
	}()

	go func() {
		receiveWrites(queue)
		wgReceive.Done()
	}()

	wg.Add(2)
	wgReceive.Add(1)

	go func() {
		wg.Wait()
		close(queue)
	}()

	wgReceive.Wait()

	fmt.Println("length:", len(anObj.getRecords()))
	fmt.Println("final:", anObj.getRecords())
}

func runWrites(start, end int, q chan int) {
	for i := start; i < end; i++ {
		//anObj.add(i)
		q <- i
	}
}

func receiveWrites(q chan int) {
	// for  {
	// 	value := <- q
	// 	anObj.add(value)
	// }
	for value := range q {
		anObj.add(value)
	}
}
