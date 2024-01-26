package main

import "fmt"

func main() {
	// START OMIT
	for i, r := range "Hello, 世界" {
		fmt.Printf("%2d: %q\n", i, r)
	}
	// END OMIT
}
