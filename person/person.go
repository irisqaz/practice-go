package person

import "fmt"

// Person type
type Person struct {
	fname string
	lname string
	age   string
}

// person.New()

// New returns an instance of a person
func New(f, l, a string) *Person {
	return &Person{f, l, a}
}

// FullName returns a string representation of the person's name
func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.fname, p.lname)
}

// Age returns the person's age as a string
func (p Person) Age() string {
	return fmt.Sprintf("%s years old", p.age)
}
