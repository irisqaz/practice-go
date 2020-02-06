package main

import (
	"fmt"
	"strconv"

	"github.com/irisqaz/practice-go/test"
)

// JoinInts returns the string representation of the joined positive numbers
// Ex: 81, 96 -> "8196"
func JoinInts(n1, n2 int) string {
	// try different solutions with
	// fmt.Sprintf
	// strconv.Itoa

	//return fmt.Sprintf("%v%v", n1, n2)
	return strconv.Itoa(n1) + strconv.Itoa(n2)
}

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(JoinInts)))

	for i := 0; i < 5; i++ {

		n1 := test.NextInt(0, 100)
		n2 := test.NextInt(0, 100)

		got := JoinInts(n1, n2)
		want := solution(n1, n2)
		isCorrect := got == want

		printResult(n1, n2, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(n1, n2 int, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%v, %v)", n1, n2)
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

func solution(n1, n2 int) string {

	//return fmt.Sprint(num)
	return fmt.Sprintf("%d%d", n1, n2)
}
