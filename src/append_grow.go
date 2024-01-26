package main

import "fmt"

func main() {
	// START OMIT
	const n = 513
	const tmpl = "len: %4d; cap: %4d\n"
	var s []int
	fmt.Printf(tmpl, len(s), cap(s))
	for i := 0; i < n; i++ {
		s = append(s, i) // HL
		fmt.Printf(tmpl, len(s), cap(s))
	}
	// END OMIT
}
