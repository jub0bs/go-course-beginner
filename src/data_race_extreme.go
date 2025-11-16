package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	const n = 2048
	var count int
	var wg sync.WaitGroup
	wg.Add(n)
	for range n {
		go func() {
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()
	fmt.Println(count) // not guaranteed to print the value of n :((((
}

// END OMIT
