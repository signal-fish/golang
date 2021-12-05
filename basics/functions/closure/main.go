package main

import "fmt"

func addTen() func(int) int {
	var n = 10
	return func(x int) int {
		n += x
		return n
	}
}

func main() {
	f := addTen()
	fmt.Println("f(1) =", f(1))
}
