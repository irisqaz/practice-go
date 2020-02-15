package main

import "fmt"

// Contact is information about a person
type Contact struct {
	fName string
	lName string
	phone string
}

// NewContact returns a pointer to a Contact
// Ex: "Jose Doe 123-456-7890" -> &Request{"Jose", "Doe", "123-456-7890"}
func NewContact(str string) *Contact {
	var fn, ln, ph string
	fmt.Sscanf(str, "%s %s %s", &fn, &ln, &ph)
	return &Contact{fn, ln, ph}
}
