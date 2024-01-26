package main

import (
	"fmt"
)

// START OMIT
func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 4; i++ {
			ch <- i
		}
		// close(ch) // HL
	}()
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("the end") // never reached! // HL
}

// END OMIT
