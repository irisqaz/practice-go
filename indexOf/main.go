package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

// IndexOf returns the index location of val if it is in nums
// otherwise it returns -1
func IndexOf(nums []int, val int) int {

	return 0
}
func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(IndexOf)))
	for i := 1; i <= 5; i++ {
		N := test.NextInt(0, 6)
		input := test.NextInts(0, 100, N)
		val := test.NextInt(0, 100)
		if N > 0 && N%2 == 0 {
			mid := len(input) / 2
			input[mid] = val
		}

		got := IndexOf(input, val)
		want := solution(input, val)
		isCorrect := got == want

		printResult(input, val, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(input []int, val int, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%v, %v)", input, val)
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

func solution(nums []int, val int) int {
	r := -1
	for i, elem := range nums {
		if val == elem {
			r = i
			break
		}
	}
	return r
}
