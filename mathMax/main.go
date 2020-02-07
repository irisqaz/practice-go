package main

import (
	"fmt"
	"math"

	"github.com/irisqaz/practice-go/test"
)

// MathMax uses math.Max to return the greater of the two given integers
// Ex: 36, 22 -> 36
func MathMax(val1, val2 int) int {
	// hint: convert int to float64 then float64 to int

	return val1
}

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(43, test.Signature(MathMax)))
	for i := 1; i <= 5; i++ {
		val1 := test.NextInt(0, 100)
		val2 := test.NextInt(0, 100)

		got := MathMax(val1, val2)
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

func solution(val1, val2 int) int {

	return int(math.Max(float64(val1), float64(val2)))
}
