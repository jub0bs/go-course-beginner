package main

import (
	"fmt"
	"sync"
)

func main() {
	// START OMIT
	const n = 1024
	var (
		count int
		mu    sync.Mutex // HL
	)
	var wg sync.WaitGroup
	wg.Add(n)
	for range n {
		go func() {
			defer wg.Done()
			mu.Lock()   // HL
			count++     // critical section // HL
			mu.Unlock() // HL
		}()
	}
	wg.Wait()
	fmt.Println(count) // exactly n; no less, no more
	// END OMIT
}
