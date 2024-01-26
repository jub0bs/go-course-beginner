package main

import "fmt"

// START OMIT
type Counter uint

func (c *Counter) Incr() { // HL
	*c++
}

func main() {
	var c Counter = 41
	c.Incr()       // HL
	fmt.Println(c) // 42
}

// END OMIT
