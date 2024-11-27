package main

import (
	"fmt"
)

// START OMIT
func main() {
	ch := make(chan int)
	go func() {
		for i := range 8 {
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
