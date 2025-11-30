package main

import "fmt"

// START OMIT
func main() {
	var (
		out = make(chan int, 1) // HL
		in  chan int            // HL
	)
	select {
	case out <- 42: // non-blocking // HL
		fmt.Println("send case")
	case v := <-in: // blocking // HL
		fmt.Println("receive case:", v)
	}
}

// END OMIT
