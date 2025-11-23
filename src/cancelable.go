package main

import (
	"fmt"
	"sync"
	"time"
)

// START1 OMIT
func cancelable(done <-chan struct{}) {
	for t := range time.Tick(500 * time.Millisecond) {
		select {
		case <-done: // HL
			return // HL
		default:
			fmt.Println(t)
		}
	}
}

// END1 OMIT

// START2 OMIT
func main() {
	done := make(chan struct{}) // HL
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		cancelable(done) // HL
	}()
	time.Sleep(5 * time.Second) // simulate more work
	close(done)                 // HL
	wg.Wait()
}

// END2 OMIT
