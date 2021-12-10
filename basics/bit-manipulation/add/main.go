package main

import "fmt"

func add(a, b int64) int64 {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}

func main() {
	fmt.Println("add(1, 2) =", add(1, 2))
	fmt.Println("add(11, 22) =", add(11, 22))
	fmt.Println("add(111, 222) =", add(111, 222))
	fmt.Println("add(1111, 2222) =", add(1111, 2222))
	fmt.Println("add(-1, -2) =", add(-1, -2))
}
