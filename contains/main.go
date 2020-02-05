package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

// Contains returns true if val is one of the elements in nums
// Ex: [12, 3, 5], 12 -> true
//     [2, 4, 10], 5  -> false
func Contains(nums []int, val int) bool {

	return true
}

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(Contains)))
	for i := 1; i <= 5; i++ {
		N := test.NextInt(0, 5)
		input := test.NextInts(0, 100, N)
		val := test.NextInt(0, 100)
		if N > 0 && N%2 == 0 {
			mid := len(input) / 2
			input[mid] = val
		}

		got := Contains(input, val)
		want := solution(input, val)
		isCorrect := test.Equal(got, want)

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

func solution(nums []int, val int) bool {
	r := false
	last := len(nums) - 1
	for i := 0; i <= last; i++ {
		if nums[i] == val {
			r = true
			break
		}
	}

	return r
}
