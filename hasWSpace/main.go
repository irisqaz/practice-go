package main

import (
	"fmt"
	"unicode"

	"github.com/irisqaz/practice-go/test"
)

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(HasWSpace)))

	for i := 0; i < 5; i++ {

		input := test.NextString()

		got := HasWSpace(input)
		want := solution(input)
		isCorrect := got == want

		printResult(input, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(w1 string, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%q)", w1)
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

func solution(str string) bool {
	r := false
	for _, ch := range str {
		if unicode.IsSpace(ch) {
			r = true
			break
		}
	}
	return r
}
