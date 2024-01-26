package main

import (
	"fmt"
)

// START OMIT
func main() {
	ch := make(chan int)
	go send(ch, 42) // HL
	fmt.Println(<-ch)
}

func send(ch chan int, n int) { // HL
	ch <- n
}

// END OMIT
