package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

// Max returns the greater of the two given numbers
// Ex: 36, 22 -> 36
func Max(val1, val2 int) int {

	return val1
}

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(43, test.Signature(Max)))
	for i := 1; i <= 5; i++ {
		val1 := test.NextInt(0, 100)
		val2 := test.NextInt(0, 100)

		got := Max(val1, val2)
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
	max := val1
	if val2 > max {
		max = val2
	}
	return max
}
