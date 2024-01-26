package main

// START OMIT
import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	go func() {
		wg.Add(1) // incorrect: concurrent with call to wg.Wait // HL
		defer wg.Done()
		fmt.Println("Grinding coffee beans...")
	}()
	wg.Wait() // HL
}

// END OMIT
