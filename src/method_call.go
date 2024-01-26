package main

import "fmt"

// START OMIT
type Counter uint

func (c Counter) Incr() Counter {
	return c + 1
}

func main() {
	var c Counter = 41
	c = c.Incr() // HL
	fmt.Println(c)
}

// END OMIT
