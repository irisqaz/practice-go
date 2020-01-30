package main

import "fmt"

func main() {
	//a := 0
	//b := 1
	a, b := 0, 1
	fmt.Println(a, b)

	a, b = b, a
	fmt.Println(a, b)

	arr := []int{a, b}
	fmt.Println(arr)
	r := swap(arr)
	fmt.Println(r)

	arr1 := []int{2, 4, 1, 9}
	fmt.Println(arr1)
	r1 := shiftLeft(arr1)
	fmt.Println(r1)

	arr2 := []int{2, 4, 1, 9}
	fmt.Println(arr2)
	r2 := shiftRight(arr2)
	fmt.Println(r2)
}

func swap(arr []int) []int {
	arr[0], arr[1] = arr[1], arr[0]
	return arr
}

func shiftLeft(arr []int) []int {
	for i := 0; i <= len(arr)-2; i++ {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}
	return arr
}

func shiftRight(arr []int) []int {
	for i := len(arr) - 1; i >= 1; i-- {
		arr[i], arr[i-1] = arr[i-1], arr[i]
	}
	return arr
}

// func shiftLeft(arr []int) []int {
// 	head := arr[0:1]
// 	tail := arr[1:]
// 	r := append(tail, head...)
// 	return r
// }
