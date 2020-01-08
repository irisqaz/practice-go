package main

import "fmt"

type mySlice []int

func main() {
	var arr mySlice = make(mySlice, 3, 10)
	arr[0] = 7
	fmt.Printf("type: %T value: %v length: %d capacity: %d", arr, arr, len(arr), cap(arr))
	fmt.Println()
	fmt.Println(arr[0])
}
