package main

import (
	"errors"
	"fmt"
	"time"
)

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(errors.New("ERROR OCCURED!!!"))
		}
	}()
	n1, n2 := 1, 0
	res := n1 / n2
	fmt.Println("res =", res)
}

func main() {
	test()
	for {
		fmt.Println("main......")
		time.Sleep(time.Second)
	}
}
