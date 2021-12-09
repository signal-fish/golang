package main

import "fmt"

func main() {
	number := 589
	squareChan := make(chan int)
	cubeChan := make(chan int)
	go calcSquares(number, squareChan)
	go calcCubes(number, cubeChan)
	squares, cubes := <-squareChan, <-cubeChan
	fmt.Println("Final output =", squares+cubes)
}

/*
1. Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel.
*/

func digits(number int, dchan chan int) {
	for number != 0 {
		digit := number % 10
		dchan <- digit
		number /= 10
	}
	close(dchan)
}

func calcSquares(number int, squareChan chan int) {
	sum := 0
	dchan := make(chan int)
	go digits(number, dchan)
	for digit := range dchan {
		sum += digit * digit
	}
	squareChan <- sum
}

func calcCubes(number int, cubeChan chan int) {
	sum := 0
	dchan := make(chan int)
	go digits(number, dchan)
	for digit := range dchan {
		sum += digit * digit * digit
	}
	cubeChan <- sum
}
