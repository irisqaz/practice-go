package main

import "fmt"

type mySlice []int

func main() {
	var arr mySlice = make(mySlice, 3, 10)
	arr[0] = 7
	fmt.Printf("type: %T value: %v length: %d capacity: %d", arr, arr, len(arr), cap(arr))
	fmt.Println()
	fmt.Println(arr[0])

	fmt.Println("Iterative empty")
	arr1 := []int{}
	travIter(arr1)
	fmt.Println()

	fmt.Println("Iterative")
	arr1 = []int{11, 33, 44}
	travIter(arr1)
	fmt.Println()

	fmt.Println("Iterative2 empty")
	arr4 := []int{}
	travIter2(arr4)
	fmt.Println()

	fmt.Println("Iterative2")
	arr4 = []int{11, 33, 44}
	travIter2(arr4)
	fmt.Println()

	fmt.Println("Recursive empty")
	arr2 := []int{}
	travIter(arr2)
	fmt.Println()

	fmt.Println("Recursive")
	arr2 = []int{11, 33, 44}
	travIter(arr2)
	fmt.Println()

	fmt.Println("Recursive2 empty")
	arr3 := []int{}
	travRec2(arr3)
	fmt.Println()

	fmt.Println("Recursive2")
	arr3 = []int{11, 33, 44}
	travRec2(arr3)
	fmt.Println()
}

func travIter(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}

func travIter2(arr []int) {
	for _, val := range arr {
		fmt.Print(val, " ")
	}
}

func travRec(arr []int) {
	arrRec(arr, 0)
}

func arrRec(arr []int, i int) {
	if i >= 0 && i < len(arr) {
		fmt.Print(arr[i], " ")
		arrRec(arr, i+1)
	}
}

func travRec2(arr []int) {
	if len(arr) > 0 {
		fmt.Print(arr[0], " ")
		travRec2(arr[1:])
	}
}
