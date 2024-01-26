package main

import "fmt"

// START OMIT
func main() {
	a := []int{1, 2}
	updateFirst(a, 0) // HL
	fmt.Println(a)
}

func updateFirst(a []int, i int) {
	a[0] = i // no length check (for brevity)
}

// END OMIT
