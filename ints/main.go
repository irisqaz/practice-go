package main

import (
	"fmt"
)

type myType int8

func max(x myType, y myType) myType {
	if x > y {
		return x
	}
	return y
}

func (m myType) String() string {
	return string(m)
}

func main() {
	myInt := myType(49)
	myInt = myInt + 1
	//myInt = myInt + 8
	//b := myInt
	//b = 100
	fmt.Println()
	fmt.Printf("%T %v %b %x %c", myInt, myInt, myInt, myInt, myInt)
	fmt.Println()
	//fmt.Printf("%T %v %b %x", b, b, b, b)
	fmt.Println()
	fmt.Println(max(myInt, myType(5)))

}
