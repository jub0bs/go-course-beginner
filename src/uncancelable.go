package main

import (
	"fmt"
	"time"
)

// START1 OMIT
func uncancelable() {
	for {
		fmt.Println(t)
		time.Sleep(500 * time.Millisecond)
	}
}

// END1 OMIT

// START2 OMIT
func main() {
	go uncancelable()
	time.Sleep(5 * time.Second) // simulate more work
}

// END2 OMIT
