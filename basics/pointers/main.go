package main

import "fmt"

/*
1. A pointer is a variable which stores the memory address of another variable.
*/

func change(val *int) {
	*val = 999
}

func main() {
	a := 666
	fmt.Println("value of a before function call is", a)
	b := &a
	change(b)
	fmt.Println("value of a after function call is", a)
}
