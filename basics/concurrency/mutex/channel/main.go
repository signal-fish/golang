package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count = 0
	ch := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(10)

	start := time.Now()
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				ch <- true
				count++
				<-ch
			}
		}()
	}
	wg.Wait()
	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("Time spent %v s\n", diff)
	fmt.Println("count =", count)
}
