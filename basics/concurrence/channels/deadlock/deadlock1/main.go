package main

import "fmt"

/*
1. If a Goroutine is sending data on a channel, then it is expected that "some other Goroutine" should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.
*/

func main() {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
}
