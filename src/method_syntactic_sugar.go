package main

// START OMIT
type Counter uint

func (c *Counter) Incr() { *c++ } // HL

func main() {
	var c Counter
	(&c).Incr() // legal, but unnecessary
	c.Incr()    // syntactic sugar! // HL
}

// END OMIT
