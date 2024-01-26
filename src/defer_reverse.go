package main

import "fmt"

// START OMIT
func main() {
	defer fmt.Println("foo")
	defer fmt.Println("bar")
	defer fmt.Println("baz")
}

//END OMIT
