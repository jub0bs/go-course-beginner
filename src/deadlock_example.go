package main

import (
	"fmt"
)

// START OMIT
func main() {
	ch := make(chan int)
	ch <- 42
	fmt.Println(<-ch)
}

// END OMIT
