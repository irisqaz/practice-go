package main

import "fmt"

func main() {
	var arr []int = make([]int, 3, 10)
	//arr = append(arr, 8)
	fmt.Printf("type: %T value: %v length: %d capacity: %d", arr, arr, len(arr), cap(arr))
	fmt.Println()
}
