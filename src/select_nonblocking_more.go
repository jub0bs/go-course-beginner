package main

import "fmt"

// START OMIT
func main() {
	var (
		out = make(chan int, 1) // HL
		in  = make(chan int, 1) // HL
	)
	in <- 42 // HL
	select {
	case out <- 42: // non-blocking // HL
		fmt.Println("send case")
		case v := <-in: // non-blocking // HL
		fmt.Println("receive case:", v)
	}
}

// END OMIT
