package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

// InEnds returns true if val is the first and last element of nums
// Ex: [12, 3, 5], 12  -> false
//     [10, 4, 10], 10 -> true
func InEnds(nums []int, val int) bool {

	return true
}

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(InEnds)))
	for i := 1; i <= 5; i++ {
		N := test.NextInt(0, 5)
		input := test.NextInts(0, 100, N)
		val := test.NextInt(0, 100)
		if N > 0 && N%2 == 0 {
			input[0] = val
			input[len(input)-1] = val
		}

		got := InEnds(input, val)
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
	/*
		r := false
		length := len(nums)
		if length > 0 {
			first := nums[0]
			last := nums[length-1]
			r = val == first && val == last
		}
		return r
	*/
	return len(nums) > 0 && val == nums[0] && val == nums[len(nums)-1]
}
