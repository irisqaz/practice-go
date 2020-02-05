package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

// IntToFloat converts an integer to a float
// Ex: 45 -> 45
func IntToFloat(num int) float64 {
	// use Type conversion
	// T(v) converts the value v to the type T

	return 0.0
}

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(IntToFloat)))

	for i := 0; i < 5; i++ {

		input := test.NextInt(-20, 20)

		got := IntToFloat(input)
		want := solution(input)
		isCorrect := got == want

		printResult(input, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(input int, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%v)", input)
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

func solution(num int) float64 {

	return float64(num)
}
