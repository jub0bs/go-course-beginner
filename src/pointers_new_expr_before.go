package main

import (
	"fmt"
	"strconv"
)

func main() {
	// START OMIT
	p := &strconv.Itoa(42) // compilation error
	fmt.Println(p)
	// END OMIT
}

