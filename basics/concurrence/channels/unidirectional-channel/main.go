package main

import "fmt"

/*
1. It is possible to convert a bidirectional channel to a send only or receive only channel but bot the vice versa.
*/

func main() {
	intChan := make(chan int)
	go sendData(intChan)
	fmt.Println(<-intChan)
}

func sendData(sendch chan<- int) {
	sendch <- 1
}
