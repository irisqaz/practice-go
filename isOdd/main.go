package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

// IsOdd returns true if val1 is odd
func IsOdd(val1 int) bool {

	return true
}

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(42, test.Signature(IsOdd)))
	for i := 1; i <= 5; i++ {
		val1 := test.NextInt(0, 100)

		got := IsOdd(val1)
		want := solution(val1)
		isCorrect := got == want

		printResult(val1, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(val1, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%v)", val1)
	strIn = test.Rjustified(25, strIn)
	strWant := test.Ljustified(20, want)
	strGot := test.Ljustified(10, got)

	if isCorrect {
		strGot = test.Green(strGot)
	} else {
		strGot = test.Red(strGot)
		strGot += fmt.Sprintf("expected: %v", test.Green(strWant))
	}

	fmt.Print(strIn, "  ->  ", strGot)
	fmt.Println()
}

func solution(val1 int) bool {

	return val1%2 != 0
}
