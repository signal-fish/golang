package main

import (
	"fmt"
	"reflect"
)

func main() {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	m1[1] = 10
	m2[1] = 10

	eq := reflect.DeepEqual(m1, m2)
	if eq {
		fmt.Println("m1 == m2")
		return
	}
	fmt.Println("m1 != m2")
}
