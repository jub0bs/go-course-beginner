package main

import (
	"context"
	"fmt"
	"time"
)

// START OMIT
func main() {
	ctx, cancel := context.WithCancel(context.Background()) // HL
	go printInts(ctx)
	time.Sleep(1 * time.Second)
	cancel() // HL
	// do more work
}

func printInts(ctx context.Context) {
	var i int
	for {
		select {
		case <-ctx.Done(): // HL
			return
		default:
			fmt.Println(i)
			i++
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// END OMIT
