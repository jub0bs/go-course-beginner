package main

import "fmt"

// START OMIT
func main() {
	a := []int{1, 2}
	updateFirst(a, 0) // HL
	fmt.Println(a)
}

func updateFirst(a []int, i int) {
	if len(a) == 0 {
		return
	}
	a[0] = i
}

// END OMIT
