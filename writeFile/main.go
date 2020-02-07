package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/irisqaz/practice-go/test"
)

// WriteFile writes the content to file named fName
// and returns the number of bytes written and error if any
func WriteFile(content, fName string) (int, error) {

	return 0, errors.New("error")
}

func main() {

	fmt.Println()
	fmt.Println(test.Signature(WriteFile))

	in1 := "This is the content\nthis is the second line\n"
	in2 := "write.txt"
	got, gotErr := WriteFile(in1, in2)
	want, wantErr := solution(in1, in2)
	isCorrect := gotErr == nil && got == want

	printResult(in1, in2, got, gotErr, want, wantErr, isCorrect)
	os.Remove(in2)
	fmt.Println()
}
func printResult(in1, in2 string, got, gotErr, want, wantErr interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%s, %s)", "...", in2)
	strIn = test.Rjustified(25, strIn)
	strWant := fmt.Sprintf("(%v, %v)", want, wantErr)
	strWant = test.Ljustified(20, strWant)
	strGot := fmt.Sprintf("(%v, %v)", got, gotErr)
	strGot = test.Ljustified(15, strGot)

	if isCorrect {
		strGot = test.Green(strGot)
	} else {
		strGot = test.Red(strGot)
		strGot += fmt.Sprintf("expected: %v", test.Green(strWant))
	}

	fmt.Print(strIn, "  ->  ", strGot)
	fmt.Println()
}

func solution(content, fName string) (int, error) {
	f, err := os.Create(fName)
	defer f.Close()
	if err != nil {
		return 0, err
	}
	return f.WriteString(content)
}
