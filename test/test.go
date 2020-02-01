package test

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// NextInt is the next random integer between a and b
func NextInt(a, b int) int {
	rand.Seed(time.Now().UnixNano())
	return a + rand.Intn(b-a+1)
}

// NextInts is a slice of length size with random integers betwee a and b
func NextInts(a, b, size int) []int {

	ints := make([]int, size)

	for i := range ints {
		ints[i] = NextInt(a, b)
	}

	return ints

}

// Equal compares the two arguments for equality and returns true when equal
func Equal(got, want interface{}) bool {
	return reflect.DeepEqual(got, want)
}

// Red converts a string to display as a red string
func Red(arg string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", arg)
}

// Green converts a string to display as a green string
func Green(arg string) string {
	return fmt.Sprintf("\033[1;32m%s\033[0m", arg)
}

// Rjustified converts arg to its string representation and justifed to the right by size spaces
func Rjustified(size int, arg interface{}) string {
	strArg := fmt.Sprintf("%v", arg)
	strSize := fmt.Sprint(size)
	format := "%" + strSize + "s"
	return fmt.Sprintf(format, strArg)
}

// Ljustified converts arg to its string representation and justifed to the left by size spaces
func Ljustified(size int, arg interface{}) string {
	strArg := fmt.Sprintf("%v", arg)
	strSize := fmt.Sprint(size)
	format := "%-" + strSize + "s"
	return fmt.Sprintf(format, strArg)
}
