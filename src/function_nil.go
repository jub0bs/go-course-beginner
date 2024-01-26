package main

import (
	"fmt"
)

func main() {
	// START OMIT
	var cleanup func()
	fmt.Println(cleanup == nil) // true
	cleanup()                   // panic!
	// END OMIT
}
