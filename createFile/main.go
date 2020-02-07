package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/irisqaz/practice-go/test"
)

// Touch creates a file named fName
// otherwise returns an error
func Touch(fName string) error {

	return errors.New("error")
}

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(40, test.Signature(Touch)))

	input := "touch.txt"
	var want error
	got := Touch(input)
	isCorrect := got == nil

	printResult(input, got, want, isCorrect)

	fmt.Println()
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
	file, err := os.Create(fName)
	defer file.Close()
	return err
}
