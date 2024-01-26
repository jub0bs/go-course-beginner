package main

import "fmt"

func main() {
	// START OMIT
	names := []string{"foo", "bar", "baz"}
	red := names[:2] // HL
	green := red[1:] // HL
	fmt.Println(red, green)
	// END OMIT
}
