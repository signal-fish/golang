package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	n := 5
	var wg sync.WaitGroup
	for i := 1; i <= n; i++ {
		wg.Add(1)
		// Must pass the pointer of WaitGroup
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished executing")
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}
