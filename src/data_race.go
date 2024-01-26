package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	var count int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		count++ // equivalent to count = count + 1
	}()
	if count == 0 {
		fmt.Println(count)
	}
	wg.Wait()
}

// END OMIT
