package main

import (
	"fmt"
	"time"
)

/*
1. The select statement blocks until one of the send/receive operationis ready. If multiple operations are ready, one of them is chosen at random.
2. The default case in a select statement is executed when none of the other cases is ready.
*/

func process(ch chan string) {
	time.Sleep(2500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value:", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
}
