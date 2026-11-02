package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Go(func() { produce(ch) })
	wg.Go(func() { consume(ch) })
	wg.Wait()
}

func produce(ch chan<- string) {
	// <-ch // would cause a compilation error // HL
	ch <- "hello"
	close(ch)
}

func consume(ch <-chan string) {
	// ch <- "hi" // would cause a compilation error // HL
	// close(ch)  // would cause a compilation error // HL
	fmt.Println(<-ch)
}

// END OMIT
