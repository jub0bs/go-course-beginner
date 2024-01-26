package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go produce(ch, &wg)
	go consume(ch, &wg)
	wg.Wait()
}

func produce(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	// <-ch // would cause a compilation error // HL
	ch <- "hello"
	close(ch)
}

func consume(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	// ch <- "hi" // would cause a compilation error // HL
	// close(ch)  // would cause a compilation error // HL
	fmt.Println(<-ch)
}

// END OMIT
