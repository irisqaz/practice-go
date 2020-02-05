package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

// Count returns the count of val in nums
// Ex: [12, 3, 5],  12 -> 1
//     [2, 10, 10], 10 -> 2
func Count(nums []int, val int) int {

	return 0
}

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(Count)))
	for i := 1; i <= 5; i++ {
		N := test.NextInt(0, 5)
		input := test.NextInts(0, 50, N)
		val := test.NextInt(0, 50)
		if N > 0 && N%2 == 0 {
			mid := N / 2
			last := N - 1
			input[mid] = val
			input[last] = val
			input[0] = val
		}

		got := Count(input, val)
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

func solution(nums []int, val int) int {
	count := 0
	last := len(nums) - 1
	for i := 0; i <= last; i++ {
		if nums[i] == val {
			count++
		}
	}

	return count
}
