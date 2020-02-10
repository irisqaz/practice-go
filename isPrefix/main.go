package main

import (
	"fmt"
	"reflect"

	"github.com/irisqaz/practice-go/test"
)

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(IsPrefix)))
	for i := 1; i <= 5; i++ {
		N := test.NextInt(0, 7)
		input := test.NextInts(0, 12, N)
		val := test.NextInts(0, 12, N/2)
		if N > 0 && N%2 == 0 {
			val = input[0 : N/2]
		}

		got := IsPrefix(input, val)
		want := solution(input, val)
		isCorrect := got == want

		printResult(input, val, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(input []int, val []int, got, want interface{}, isCorrect bool) {
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

func solution(nums1 []int, nums2 []int) bool {

	return len(nums1) >= len(nums2) && reflect.DeepEqual(nums1[0:len(nums2)], nums2)
}
