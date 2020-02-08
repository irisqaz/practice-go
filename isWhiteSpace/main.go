package main

import (
	"fmt"
	"unicode"

	"github.com/irisqaz/practice-go/test"
)

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(42, test.Signature(IsWhiteSpace)))
	for i := 1; i <= 5; i++ {
		char := test.NextChar()

		got := IsWhiteSpace(char)
		want := solution(char)
		isCorrect := got == want

		printResult(char, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(val1, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%q)", val1)
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

func solution(char byte) bool {
	return unicode.IsSpace(rune(char))
}
