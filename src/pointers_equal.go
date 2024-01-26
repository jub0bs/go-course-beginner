package main

import "fmt"

func main() {
	// START OMIT
	i := 42
	p1 := &i
	p2 := &i
	fmt.Println(p1 == p2) // true
	// END OMIT
}
