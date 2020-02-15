package main

import (
	"fmt"
	"reflect"

	"github.com/irisqaz/practice-go/test"
)

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(NewContact)))

	input := []string{
		"Jose Doe 123-456-7890",
		"Sabrina Lee 415-123-4567",
		"Dora Explorer 916-234-5678",
	}
	for _, s := range input {

		got := NewContact(s)
		want := solution(s)
		isCorrect := reflect.DeepEqual(got, want)

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

func solution(str string) *Contact {
	var fn, ln, phone string
	fmt.Sscanf(str, "%s %s %s", &fn, &ln, &phone)
	return &Contact{fn, ln, phone}
}
