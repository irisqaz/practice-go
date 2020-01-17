package main

import (
	"fmt"
	"os"
)

var gInt int = 100

func main() {
	fmt.Println()
	fmt.Fprintln(os.Stdout)

	_ = add(5, 3)

	a := 5
	b := 3

	addP(a, &b)
	fmt.Println(b)

	var c int
	var d int

	//fmt.Fscan(os.Stdin, &c, &d)
	fmt.Println(c, d)

	e := []int{5, 6, 7}
	func1(e, func(x int) {
		fmt.Println("I got:", x)
	})

	funcG()
	gInt = 10
	funcG()
}

func funcG() {
	fmt.Println("The global is now:", gInt)
}

func func1(arr []int, f func(n2 int)) {
	for _, val := range arr {
		f(val)
	}
}

func addP(n1 int, n2 *int) {
	*n2 = n1 + *n2
}

func add(a, b int) int {
	return a + b
}
