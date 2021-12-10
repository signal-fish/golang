package main

import "fmt"

func swap(a, b *int) (int, int) {
	*a = *a ^ *b
	*b = *b ^ *a
	*a = *a ^ *b
	return *a, *b
}

func main() {
	a, b := 1, 2
	fmt.Printf("Before Swap, a = %v, b = %v\n", a, b)
	swap(&a, &b)
	fmt.Printf("After Swap, a = %v, b = %v\n", a, b)
}
