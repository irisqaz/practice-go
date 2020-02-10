package main

import (
	"fmt"
	"strings"

	"github.com/irisqaz/practice-go/test"
)

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(HasSuffix)))

	for i := 0; i < 5; i++ {

		input := test.StringOfSize(10)
		val := test.StringOfSize(3)

		if i%2 == 0 {
			val = input[len(input)-3:]
		}

		got := HasSuffix(input, val)
		want := solution(input, val)
		isCorrect := got == want

		printResult(input, val, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(w1 string, val string, got, want bool, isCorrect bool) {
	strIn := fmt.Sprintf("(%q, %q)", w1, val)
	strIn = test.Rjustified(25, strIn)
	//strWant := test.Quote(want)
	strWant := test.Ljustified(20, want)
	//strGot := test.Quote(got)
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

func solution(str1, str2 string) bool {

	return strings.HasSuffix(str1, str2)
}
