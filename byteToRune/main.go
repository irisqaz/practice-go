package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(ByteToRune)))

	for i := 0; i < 5; i++ {

		input := test.NextChar()

		got := ByteToRune(input)
		want := solution(input)
		isCorrect := got == want

		printResult(input, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(input byte, got, want rune, isCorrect bool) {
	strIn := fmt.Sprintf("(%q)", input)
	strIn = test.Rjustified(25, strIn)
	strWant := fmt.Sprintf("%q", input)
	strWant = test.Ljustified(20, strWant)
	strGot := fmt.Sprintf("%q", got)
	strGot = test.Ljustified(10, strGot)

	if isCorrect {
		strGot = test.Green(strGot)
	} else {
		strGot = test.Red(strGot)
		strGot += fmt.Sprintf("expected: %s", test.Green(strWant))
	}

	fmt.Print(strIn, "  ->  ", strGot)
	fmt.Println()
}

func solution(char byte) rune {

	return rune(char)
}
