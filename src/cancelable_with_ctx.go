package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// START1 OMIT
func cancelable(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // HL
			return // HL
		default:
			fmt.Println(t)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// END1 OMIT

// START2 OMIT
func main() {
	ctx, cancel := context.WithCancel(context.Background()) // HL
	defer cancel()                                          // HL
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		cancelable(ctx) // HL
	}()
	time.Sleep(5 * time.Second) // simulate more work
	cancel()                    // HL
	wg.Wait()
}

// END2 OMIT
