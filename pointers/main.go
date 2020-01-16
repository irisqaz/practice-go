package main

import "fmt"

func main() {
	int1 := 6
	var pInt1 *int = &int1
	fmt.Printf("type: %T value: %v \n", int1, int1)
	fmt.Printf("type: %T value: %v \n", pInt1, pInt1)
	*pInt1 = 10
	fmt.Printf("type: %T value: %v \n", int1, int1)
	fmt.Printf("type: %T value: %v \n", pInt1, pInt1)
}
