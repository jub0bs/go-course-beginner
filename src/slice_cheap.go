package main

import "fmt"

// START OMIT
func main() {
	a := []int{1, 2}
	fmt.Println(&a[0])
	p := addressOfFirst(a)
	fmt.Println(p)
	fmt.Println(&a[0] == p) // true // HL
}

func addressOfFirst(a []int) *int {
	if len(a) == 0 {
		return nil
	}
	return &a[0]
}

// END OMIT
