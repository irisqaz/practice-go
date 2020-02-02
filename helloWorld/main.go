package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

// HelloWorld returns a string containing "Hello World!"
func HelloWorld() string {

	return "?"
}
func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(HelloWorld)))

	got := HelloWorld()
	want := solution()
	isCorrect := got == want

	printResult(got, want, isCorrect)
	fmt.Println()
}
func printResult(got, want interface{}, isCorrect bool) {
	strIn := "()"
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

func solution() string {

	return "Hello World!"
}
