package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/irisqaz/practice-go/test"
)

// OpenReading opens an existing file named fName
// otherwise returns an error
func OpenReading(fName string) error {

	return errors.New("error")
}

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(40, test.Signature(OpenReading)))

	input := "touch.txt"
	var want error
	got := OpenReading(input)
	isCorrect := got == nil

	printResult(input, got, want, isCorrect)
	fmt.Println()

	if isCorrect {
		os.Remove(input)
	}
}
func printResult(fName string, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%s)", fName)
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

func solution(fName string) error {
	file, err := os.Open(fName)
	defer file.Close()
	return err
}
