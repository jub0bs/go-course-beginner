package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	fmt.Println(replicatedFetch("Strasbourg", "Paris"))
}

func replicatedFetch(dc1, dc2 string) string {
	ch := make(chan string)
	// for brevity, I've omitted the impl. of function fetch (of type func(string) string)
	go func() { ch <- fetch(dc1) }() // HL
	go func() { ch <- fetch(dc2) }() // HL
	return <-ch                      // HL
}

// END OMIT

func fetch(dc string) string {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	return dc // simply return the name of the datacentre
}
