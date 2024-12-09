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

func replicatedFetch(dcs ...string) string {
	ch := make(chan string)
	for _, dc := range dcs {
		go func() { ch <- fetch(dc) }() // fetch is just some func(string) string // HL
	}
	return <-ch // HL
}

// END OMIT

func fetch(dc string) string {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	return dc // simply return the name of the datacentre
}
