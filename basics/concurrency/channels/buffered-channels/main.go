package main

import (
	"fmt"
	"time"
)

/*
1. Sends to a buffered channel are blocked only when the buffer is full and receives from a buffered channel are blocked only when the buffer is empty.
*/

func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("Read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}

func write(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Println("Successfully wrote", i, "to ch")
	}
	close(ch)
}
