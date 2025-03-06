package main

import "fmt"

// START OMIT
func main() {
	var (
		out chan int // HL
		in  chan int // HL
	)
	select {
	case out <- 42: // blocking // HL
		fmt.Println("send case")
	case <-in: // blocking // HL
		fmt.Println("receive case")
	default: // HL
		fmt.Println("default case")
	}
}

// END OMIT
