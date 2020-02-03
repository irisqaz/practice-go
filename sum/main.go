package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

// Sum returns the sum of all the values in nums
func Sum(nums []int) int {

	return 0
}
func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(Sum)))
	for i := 1; i <= 5; i++ {
		N := test.NextInt(0, 5)
		input := test.NextInts(0, 80, N)

		got := Sum(input)
		want := solution(input)
		isCorrect := got == want

		printResult(input, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(input []int, got, want interface{}, isCorrect bool) {
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

func solution(nums []int) int {
	sum := 0
	last := len(nums) - 1
	for i := 0; i <= last; i++ {
		sum += nums[i]
	}

	return sum
}
