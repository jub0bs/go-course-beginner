package main

import "fmt"

func main() {
	// START OMIT
	var (
		m1 map[int]string
		m2 map[[]int]string // compilation error!
	)
	// END OMIT
	fmt.Println(m1, m2)
}
