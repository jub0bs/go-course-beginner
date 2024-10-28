package main

import "fmt"

// START OMIT
func main() {
	nums := []int{1, 2}
	updateFirst(nums, 0) // HL
	fmt.Println(nums)
}

func updateFirst(s []int, i int) {
	if len(s) == 0 {
		return
	}
	s[0] = i // HL
}

// END OMIT
