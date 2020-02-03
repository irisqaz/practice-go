package main

import "fmt"

func main() {
	freq := map[string]int{}
	freq["apple"] = 1
	freq["apple"]++
	_, in := freq["orange"]
	if in {
		freq["orange"]++
	} else {
		freq["orange"] = 1
	}
	fmt.Println(freq)
}
