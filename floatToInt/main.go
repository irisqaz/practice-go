package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

// FloatToInt converts a float to an integer
// Ex: 3.45 -> 3
func FloatToInt(num float64) int {
	// use Type conversion
	// T(v) converts the value v to the type T

	return 0
}

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(FloatToInt)))

	for i := 0; i < 5; i++ {

		input := test.NextFloat(10)

		got := FloatToInt(input)
		want := solution(input)
		isCorrect := got == want

		printResult(input, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(input float64, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%.2f)", input)
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

func solution(num float64) int {

	return int(num)
}
