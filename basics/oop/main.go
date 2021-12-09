package main

import "golang-basics/basics/oop/employee"

func main() {
	e := employee.New("Signal", "Fish", 25)
	e.Describe()
}
