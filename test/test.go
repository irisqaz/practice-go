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
var text1 string
var text2 string

func init() {

	text1 = "firm leg\tfish mind\nband bark lung rung role\tsoul dull fall lay news end lock\tbaby mug tick show flu\nrow glue yard"

	words = strings.Fields(text1)
}

// NextChar returns a random character
func NextChar() byte {
	next := NextInt(0, len(text1)-1)
	return text1[next]
}

// NextWord returns a random word
func NextWord() string {
	next := NextInt(0, len(words)-1)
	return words[next]
}

// NextString returns a string of random size
func NextString() string {
	last := len(text1) - 1
	start := NextInt(0, last/2)
	end := NextInt(start, last)
	if end-start > 10 {
		end = start + 10
	}
	return text1[start:end]
}

// StringOfSize returns a random string of length 'size' or less
func StringOfSize(size int) string {
	last := len(text1) - 1
	start := NextInt(0, last/2)
	end := start + size
	if end > last {
		end = last
	}
	return text1[start:end]
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

// NextFloat is the next random float64 between a and b
func NextFloat(a, b float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return a + rand.Float64()*(b-a)
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

// Quote returns string str in quotes
func Quote(str string) string {
	return fmt.Sprintf("%q", str)
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
