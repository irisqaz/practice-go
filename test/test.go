package test

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"time"
)

var words []string

func init() {

	var text = `firm leg fish mind band bark lung rung role soul dull fall lay 
                news end lock baby mug tick show flu row glue yard`

	words = strings.Fields(text)
}

// NextWord returns a random word
func NextWord() string {
	next := NextInt(0, len(words)-1)
	return words[next]
}

// NextInt is the next random integer between a and b
func NextInt(a, b int) int {
	rand.Seed(time.Now().UnixNano())
	return a + rand.Intn(b-a+1)
}

// NextInts is a slice of length size with random integers between a and b
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

// Yellow converts a string to display as a yellow string
func Yellow(arg string) string {
	return fmt.Sprintf("\033[1;33m%s\033[0m", arg)
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

// Signature returns the string signature of function f
func Signature(f interface{}) string {
	sign := fmt.Sprintf("%T", f)
	sign = strings.ReplaceAll(sign, "func", "")
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	name = strings.Split(name, ".")[1]
	name = Yellow(name)
	return name + sign
}
