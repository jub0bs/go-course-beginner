package main

import (
	"fmt"
)

// START OMIT
func main() {
	ca := make(chan int)
	cb := ca // copy // HL
	go func() {
		cb <- 42
	}()
	fmt.Println(<-ca)
	close(ca)
	v, ok := <-cb
	fmt.Println(v, ok)
}

// END OMIT
