package main

import "fmt"

func main() {
	fmt.Println(mult(6, 0))
	fmt.Println(pow(2, 3))
	fmt.Println(powMult(5, 0))
	fmt.Println(multRec(10, 1))
	fmt.Println(powRec(2, 3))
}

func mult(n1 int, n2 int) int {

	r := 0

	for i := 1; i <= n2; i++ {
		r = n1 + r
	}

	return r
}

func multRec(n1 int, n2 int) int {

	r := 0

	if n2 >= 1 {
		r = n1 + multRec(n1, n2-1)
	}

	return r
}

func powMult(n1 int, n2 int) int {

	r := 1

	for i := 1; i <= n2; i++ {
		r = mult(r, n1)
	}

	return r
}

func pow(n1 int, n2 int) int {

	r := 1

	for i := 1; i <= n2; i++ {
		r = n1 * r
	}

	return r
}

func powRec(n1 int, n2 int) int {

	r := 1

	if n2 >= 1 {
		r = n1 * powRec(n1, n2-1)
	}

	return r
}
