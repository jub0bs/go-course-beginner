package main

import "fmt"

func main() {
	// START OMIT
	a := []int{1, 2}
	b := a // b is an "alias" of a
	fmt.Println(a, b)
	b[0] = 0 // this change affects both b and a // HL
	fmt.Println(a, b)
	// END OMIT
}
