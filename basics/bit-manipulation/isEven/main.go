package main

import "fmt"

func isEven(n int) bool {
	return n&1 == 1
}

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	for i, v := range nums {
		if isEven(v) {
			fmt.Printf("nums[%d] = %v, is an even\n", i, v)
		} else {
			fmt.Printf("nums[%d] = %v, is an old\n", i, v)
		}
	}
}
