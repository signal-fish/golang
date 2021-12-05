package main

import "fmt"

func main() {
	res1 := func(n1, n2 int) int {
		return n1 + n2
	}(1, 1)
	fmt.Println("res1 =", res1)

	add := func(n1, n2 int) int {
		return n1 + n2
	}
	res2 := add(11, 11)
	fmt.Println("res2 =", res2)
}
