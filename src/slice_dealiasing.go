package main

import "fmt"

func main() {
	// START OMIT
	blue := []int{1, 2}
	red := blue // red is an "alias" of blue
	fmt.Println(blue, red)
	red = append(red, 3) // decouples red from blue! // HL
	fmt.Println(blue, red)
	red[0] = 0
	fmt.Println(blue, red) // updating red doesn't change blue
	// END OMIT
}
