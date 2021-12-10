package main

import "fmt"

func devideByTwo(n uint) uint {
	return n >> 1
}

func main() {
	nums := [5]uint{1, 2, 3, 4, 5}
	for i, v := range nums {
		fmt.Printf("nums[%d] = %v, devided by 2 ==> %v\n", i, v, devideByTwo(v))
	}
}
