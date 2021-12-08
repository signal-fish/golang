package main

import "fmt"

type Employee struct {
	name string
	age  int
}

func main() {
	e := Employee{
		name: "Signal Fish",
		age:  25,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Signal Yu")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	e.changeAge(23)
	fmt.Printf("\nEmployee age after change: %d", e.age)
}

// Method with value receiver
func (e Employee) changeName(newName string) {
	e.name = newName
}

// Method with pointer receiver
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}
