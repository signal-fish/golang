package main

import "fmt"

/*
1. An array is a collection of elements that belong to the same type.
2. Arrays in Go are value types and not reference types.
*/

func main() {
	num := [...]int{333, 666, 999}
	fmt.Println("Before passing to changeLocal function, num =", num)
	changeLocal(num)
	fmt.Println("After passing to changeLocal function, num =", num)

	// Traverse array
	for i, v := range num {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
}

func changeLocal(num [3]int) {
	num[2] = 888
	fmt.Println("Inside changeLocal function, num =", num)
}
