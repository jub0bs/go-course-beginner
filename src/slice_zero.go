package main

import "fmt"

func main() {
	// START OMIT
	var ids []int   // nil slice // HL
	ids2 := []int{} // empty slice // HL
	fmt.Println(ids, ids2)
	fmt.Println(ids == nil, ids2 == nil)       // Don't do this.
	fmt.Println(len(ids) == 0, len(ids2) == 0) // Do this instead. // HL
	// END OMIT
}
