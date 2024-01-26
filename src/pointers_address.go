package main

import "fmt"

func main() {
	// START OMIT
	i := 42
	p := &i // HL
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	// END OMIT
}
