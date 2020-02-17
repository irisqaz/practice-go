package main

import (
	"fmt"
	"sync"
	"time"
)

type sharedObj struct {
	records []int
	key     sync.Mutex
}

func (obj *sharedObj) add(n int) {
	obj.key.Lock()
	defer obj.key.Unlock()
	obj.records = append(obj.records, n)

}

func (obj *sharedObj) getRecords() []int {
	return obj.records
}

var anObj sharedObj

func main() {
	var wg sync.WaitGroup

	go func() {
		runWrites(0, 10)
		wg.Done()
	}()

	go func() {
		runWrites(10, 20)
		wg.Done()
	}()

	wg.Add(2)
	wg.Wait()

	fmt.Println("length:", len(anObj.getRecords()))
	fmt.Println("final:", anObj.getRecords())
}

func runWrites(start, end int) {
	for i := start; i < end; i++ {
		anObj.add(i)
		time.Sleep(time.Millisecond * 1)
	}
}
