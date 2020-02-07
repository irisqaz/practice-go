package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

// IsSameOdd returns true if val1 is equal to val2 and both are Odd
func IsSameOdd(val1, val2 int) bool {

	return true
}

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(42, test.Signature(IsSameOdd)))
	for i := 1; i <= 5; i++ {
		val1 := test.NextInt(0, 100)
		val2 := test.NextInt(0, 100)
		if val1%2 != 0 {
			val2 = val1
		}

		got := IsSameOdd(val1, val2)
		want := solution(val1, val2)
		isCorrect := got == want

		printResult(val1, val2, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(val1, val2, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%v, %v)", val1, val2)
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

func solution(val1, val2 int) bool {

	return val1 == val2 && isOdd(val1)
}

func isOdd(val1 int) bool {

	return val1%2 != 0
}
