package main

import "fmt"

func main() {
	// START OMIT
	names := []string{"foo", "bar", "baz"}
	red := names[:2]
	fmt.Println(red[:cap(red)]) // HL
	// END OMIT
}
