package main

import "fmt"

func main() {
	fmt.Println(mult(6, 0))
	fmt.Println(pow(2, 3))
	fmt.Println(powMult(5, 0))
	fmt.Println(multRec(0, 5))
}

func mult(n1 int, n2 int) int {

	r := 0

	for i := 1; i <= n2; i++ {
		r = r + n1
	}

	return r
}

func multRec(n1 int, n2 int) int {

	if n2 > 1 {
		return n1 + multRec(n1, n2-1)
	}
	return 0
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
		r = r * n1
	}

	return r
}
