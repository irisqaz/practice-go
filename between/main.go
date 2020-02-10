package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(Between)))
	for i := 1; i <= 5; i++ {
		N := test.NextInt(0, 7)
		input := test.NextInts(0, 12, N)
		last := len(input) - 1
		if last < 0 {
			last = 0
		}
		from := test.NextInt(0, last/2)
		to := test.NextInt(from, last)

		got := Between(input, from, to)
		want := solution(input, from, to)
		isCorrect := test.Equal(got, want)

		printResult(input, from, to, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(input []int, from, to int, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%v, %d, %d)", input, from, to)
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

func solution(nums []int, from, to int) []int {

	return nums[from:to]
}
