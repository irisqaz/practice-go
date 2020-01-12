package main

import "fmt"

type contact = struct {
	name  string
	phone int
}

func main() {
	obj := contact{name: "jose", phone: 123}
	obj2 := contact{name: "nancy", phone: 234}

	myContacts := []contact{obj, obj2}

	fmt.Println(myContacts)

	for index := 0; index < len(myContacts); index++ {
		myContacts[index].name = "nancy"
	}

	for _, val := range myContacts {
		fmt.Println(val.name)
	}

}
