package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

// JoinStr returns the two given strings joined by a space
// Ex: "word1", "word2" -> "word1 word2"
func JoinStr(w1, w2 string) string {
	// try different solutions with
	// +
	// fmt.Sprintf

	//return w1 + " " + w2
	return fmt.Sprintf("%s %s", w1, w2)

}

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(41, test.Signature(JoinStr)))

	for i := 0; i < 5; i++ {

		w1 := test.NextWord()
		w2 := test.NextWord()

		got := JoinStr(w1, w2)
		want := solution(w1, w2)
		isCorrect := got == want

		printResult(w1, w2, got, want, isCorrect)
	}
	fmt.Println()
}
func printResult(w1, w2 string, got, want interface{}, isCorrect bool) {
	strIn := fmt.Sprintf("(%v, %v)", w1, w2)
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

func solution(w1, w2 string) string {

	//return w1 + " " + w2
	return fmt.Sprintf("%v %v", w1, w2)
}
