package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

// InFirst returns true if val is the first element of nums
func InFirst(nums []int, val int) bool {

	return true
}
func main() {

	fmt.Println()
	for i := 1; i <= 5; i++ {
		N := test.NextInt(0, 5)
		input := test.NextInts(0, 100, N)
		val := test.NextInt(0, 100)
		if N > 0 && N%2 == 0 {
			input[0] = val
		}

		got := InFirst(input, val)
		want := solution(input, val)
		isCorrect := test.Equal(got, want)

		printResult(input, val, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(input []int, val int, got, want interface{}, isCorrect bool) {

	strIn := test.Rjustified(25, input)
	strWant := test.Ljustified(20, want)
	strGot := test.Ljustified(10, got)
	strVal := " , " + test.Rjustified(3, val)

	if isCorrect {
		strGot = test.Green(strGot)
	} else {
		strGot = test.Red(strGot)
		strGot += fmt.Sprintf("expected: %v", test.Green(strWant))
	}

	fmt.Print(strIn, strVal, "  ->  ", strGot)
	fmt.Println()
}

func solution(nums []int, val int) bool {

	return len(nums) > 0 && val == nums[0]
}
