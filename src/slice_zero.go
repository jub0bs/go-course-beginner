package main

import "fmt"

func main() {
	// START OMIT
	var ids []int
	fmt.Println(ids == nil) // HL
	fmt.Println(ids)

	ids = []int{} // empty slice // HL
	fmt.Println(ids)
	// END OMIT
}
