package main

import (
	"fmt"
	"time"
)

// START1 OMIT
func uncancelable() {
	for t := range time.Tick(500 * time.Millisecond) {
		fmt.Println(t)
	}
}

// END1 OMIT

// START2 OMIT
func main() {
	go uncancelable()
	time.Sleep(5 * time.Second) // simulate more work
}

// END2 OMIT
