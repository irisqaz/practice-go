package mypkg

import (
	"fmt"
	"strings"
)

// Contains returns true when str2 is in str1
func Contains(str1 string, str2 string) bool {

	return strings.Contains(str1, str2)
}

// Prompt prompts user for a value
func Prompt(p string) string {
	var value string
	fmt.Printf("%s: ", p)
	fmt.Scanln(&value)
	return value
}
