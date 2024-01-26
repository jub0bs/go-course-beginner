package main

import "fmt"

func main() {
	// START OMIT
	var p *int      // nil
	fmt.Println(*p) // panic!
	// END OMIT
}
