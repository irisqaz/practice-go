package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(42, test.Signature(DiffFloat)))
	for i := 1; i <= 5; i++ {
		val1 := test.NextFloat(-10, 10)
		val2 := test.NextFloat(-10, 10)
		if i%2 == 0 {
			val2 = val1 - .0009
		}

		got := DiffFloat(val1, val2)
		want := solution(val1, val2)
		isCorrect := got == want

		printResult(val1, val2, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(val1, val2, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%.5f, %.5f)", val1, val2)
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

func solution(val1, val2 float64) bool {

	return abs(val1-val2) < .001
}

func abs(val float64) float64 {
	if val < 0 {
		return -val
	}
	return val
}
