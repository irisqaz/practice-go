package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(HTMLElement)))

	m := map[string]string{}
	m["p"] = "a paragraph"
	m["div"] = "a div section"
	m["h1"] = "a heading 1"
	m["body"] = "the body"
	for elem, text := range m {

		got := HTMLElement(elem, text)
		want := solution(elem, text)
		isCorrect := got == want

		printResult(elem, text, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(w1, w2 string, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%q, %q)", w1, w2)
	strIn = test.Rjustified(25, strIn)
	strWant := fmt.Sprintf("%q", want)
	strWant = test.Ljustified(20, strWant)
	strGot := fmt.Sprintf("%q", got)
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

func solution(elem, text string) string {

	return fmt.Sprintf("<%s>%s</%s>", elem, text, elem)
}
