package main

import (
	"fmt"
	"strconv"

	"github.com/irisqaz/practice-go/test"
)

// IntToStr returns the string representation of the number
func IntToStr(num int) string {
	// try different solutions with
	// fmt.Sprint
	// fmt.Sprintf
	// strconv.Itoa

	return "?"
}
func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(IntToStr)))

	for i := 0; i < 5; i++ {

		input := test.NextInt(-20, 20)

		got := IntToStr(input)
		want := solution(input)
		isCorrect := got == want

		printResult(input, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(input int, got, want interface{}, isCorrect bool) {
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

func solution(num int) string {

	//return fmt.Sprint(num)
	return strconv.Itoa(num)
}
