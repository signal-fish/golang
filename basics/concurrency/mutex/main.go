package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var count = 0

	var wg sync.WaitGroup
	wg.Add(10)

	start := time.Now()
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("Time spent %v s", diff)
	fmt.Println("count =", count)
}
