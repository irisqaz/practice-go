package mypkg

import (
	"strings"
)

// Contains returns true when str2 is in str1
func Contains(str1 string, str2 string) bool {

	return strings.Contains(str1, str2)
}
