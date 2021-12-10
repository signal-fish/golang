package main

import "fmt"

func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0
}

func main() {
	nums := [5]int{1, 2, 4, 8, 9}
	for i, v := range nums {
		if isPowerOfTwo(v) {
			fmt.Printf("nums[%d] = %v, is power of two\n", i, v)
		} else {
			fmt.Printf("nums[%d] = %v, is not power of two\n", i, v)
		}
	}
}
