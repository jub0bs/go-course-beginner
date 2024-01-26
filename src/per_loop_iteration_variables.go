package main

import (
	"fmt"
	"sync"
)

func main() {
	printTenIntsConcurrently()
}

// START OMIT
func printTenIntsConcurrently() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ { // HL
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i) // HL
		}()
	}
	wg.Wait()
}

// END OMIT
