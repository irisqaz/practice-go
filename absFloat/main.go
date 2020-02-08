package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(42, test.Signature(AbsFloat)))
	for i := 1; i <= 5; i++ {
		val := test.NextFloat(-10, 10)

		got := AbsFloat(val)
		want := solution(val)
		isCorrect := got == want

		printResult(val, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(val1, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%v)", val1)
	strIn = test.Rjustified(25, strIn)
	strWant := test.Ljustified(20, want)
	strGot := test.Ljustified(22, got)

	if isCorrect {
		strGot = test.Green(strGot)
	} else {
		strGot = test.Red(strGot)
		strGot += fmt.Sprintf("expected: %v", test.Green(strWant))
	}

	fmt.Print(strIn, "  ->  ", strGot)
	fmt.Println()
}

func solution(val float64) float64 {
	if val < 0 {
		return -val
	}
	return val
}
