package main

import (
	"fmt"
)

func main() {

	/* 	ints := []byte{72, 101, 108, 108, 111, 33}

	   	fmt.Println()
	   	fmt.Printf("%v %s", ints, ints)
	   	fmt.Println()

	   	ints[0] = 73

	   	str := string(ints)

	   	str = str + " world"

	   	str = strings.ToUpper(str)

	   	fmt.Println()
	   	fmt.Printf("%v %s", str, str)
	   	fmt.Println()

	   	for i := 0; i < len(str); i++ {
	   		fmt.Println(str[i])
	   	}

	   	for _, chr := range str {
	   		fmt.Printf("%c \n", chr)
	   	} */

	str := "hello"
	str = str + " world"
	fmt.Println(str)

	sub := str[2:]
	fmt.Printf("%v", sub)
	fmt.Println()

}
