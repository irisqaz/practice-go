package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

// Or returns the boolean 'or' operation on val1 and val2
func Or(val1, val2 bool) bool {

	return true
}

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(40, test.Signature(Or)))
	for i := 1; i <= 5; i++ {
		val1 := test.NextInt(1, 10)%2 == 0
		val2 := test.NextInt(1, 10)%2 == 0
		if i%2 == 0 {
			val1 = val2
		}
		got := Or(val1, val2)
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

func solution(val1, val2 bool) bool {

	return val1 || val2
}
