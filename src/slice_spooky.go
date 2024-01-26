package main

import "fmt"

func main() {
	// START OMIT
	blue := []int{1, 2}
	red := blue // red is an "alias" of blue
	fmt.Println(blue, red)
	red[0] = 0 // this change affects both red and blue // HL
	fmt.Println(blue, red)
	// END OMIT
}
