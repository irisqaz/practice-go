package main

import "fmt"

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
	go runWrites(0, 10)
	runWrites(10, 20)

	fmt.Println("length:", len(anObj.getRecords()))
	fmt.Println("final:", anObj.getRecords())
}

func runWrites(start, end int) {
	for i := start; i < end; i++ {
		anObj.add(i)
	}
}
