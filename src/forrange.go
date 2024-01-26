package main

import "fmt"

func main() {
	// START OMIT
	msg := "Hello, 世界"
	for i, r := range msg {
		fmt.Printf("%2d: %q\n", i, r)
	}
	// END OMIT
}
