package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(NewRequest)))

	input := []string{
		"GET /hello.html HTTP/1.1",
		"PUT /job/123 HTTP/1.1",
		"POST / HTTP/1.1",
	}
	for _, s := range input {

		got := NewRequest(s)
		want := solution(s)
		isCorrect := got == want

		printResult(s, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(w1 string, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%q)", w1)
	strIn = test.Rjustified(30, strIn)
	strWant := test.Ljustified(20, want)
	//strGot := fmt.Sprintf("%+v", got)
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

func solution(str string) *Request {
	var meth, path, prot string
	fmt.Sscanf(str, "%s %s %s", &meth, &path, &prot)
	return &Request{meth, path, prot}
}
