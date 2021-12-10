package main

import (
	"fmt"
	"unsafe"
)

/*
1. The most important property of the empty struct is that the width is zero.
2. Use of empty struct
   2-1. As a method receiver
   2-2. An empty struct channel
   2-3. As a Set data type: map[key]struct{}
*/

func main() {
	var s struct{}
	var i interface{}
	var b bool

	fmt.Println("Size of s:", unsafe.Sizeof(s))
	fmt.Println("Size of i:", unsafe.Sizeof(i))
	fmt.Println("Size of b:", unsafe.Sizeof(b))

}
