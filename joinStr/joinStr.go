package main

import "fmt"

// JoinStr returns the two given strings joined by a space
// Ex: "word1", "word2" -> "word1 word2"
func JoinStr(w1, w2 string) string {
	// try different solutions with
	// +
	// fmt.Sprintf

	//return w1 + " " + w2
	return fmt.Sprintf("%s %s", w1, w2)

}
