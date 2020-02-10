package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(Substring)))

	for i := 0; i < 5; i++ {

		input := test.NextString()
		last := len(input) - 1
		if last < 0 {
			last = 0
		}
		from := test.NextInt(0, last/2)
		to := test.NextInt(from, last)

		got := Substring(input, from, to)
		want := solution(input, from, to)
		isCorrect := got == want

		printResult(input, from, to, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(w1 string, from, to int, got, want string, isCorrect bool) {
	strIn := fmt.Sprintf("(%q, %d, %d)", w1, from, to)
	strIn = test.Rjustified(25, strIn)
	strWant := test.Quote(want)
	strWant = test.Ljustified(20, strWant)
	strGot := test.Quote(got)
	strGot = test.Ljustified(10, strGot)

	if isCorrect {
		strGot = test.Green(strGot)
	} else {
		strGot = test.Red(strGot)
		strGot += fmt.Sprintf("expected: %v", test.Green(strWant))
	}

	fmt.Print(strIn, "  ->  ", strGot)
	fmt.Println()
}

func solution(str string, from, to int) string {

	return str[from:to]
}
