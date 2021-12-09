package employee

import "fmt"

type employee struct {
	firstName string
	lastName  string
	age       int
}

func New(firstName, lastName string, age int) employee {
	e := employee{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
	return e
}

func (e employee) Describe() {
	fmt.Printf("%s %s is %d years old.\n", e.firstName, e.lastName, e.age)
}
