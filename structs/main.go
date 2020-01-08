package main

import "fmt"

type contact = struct {
	name  string
	phone int
}

func main() {
	obj := contact{name: "jose", phone: 123}
	obj2 := contact{name: "nancy", phone: 234}

	arr := []contact{obj, obj2}

	fmt.Println(arr)
}
