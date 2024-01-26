package main

import "fmt"

func main() {
	// START OMIT
	names := []string{"foo", "bar", "baz"}
	red := names[0:2] // HL
	fmt.Printf("red: %v; len %d; cap %d\n", red, len(red), cap(red))
	fmt.Println(&names[0] == &red[0] && &names[1] == &red[1])
	green := red[1:2] // HL
	fmt.Printf("green: %v; len %d; cap %d\n", green, len(green), cap(green))
	fmt.Println(&names[1] == &green[0])
	// END OMIT
}
