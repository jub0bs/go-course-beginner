package main

import "fmt"

// START OMIT
func main() {
	fmt.Println(Min(42, 999, -1, 23), Max(44, 99))
}

func Min(z int, nz ...int) int {
	f := func(x, y int) int { return min(x, y) } // HL
	return Reduce(f, z, ns...)
}

func Max(z int, nz ...int) int {
	f := func(x, y int) int { return max(x, y) } // HL
	return Reduce(f, z, ns...)
}

func Reduce(f func(acc, n int) int, z int, ns ...int) int {
	// implementation omitted
	// END OMIT
	for _, n := range ns {
		z = f(z, n)
	}
	return z
}
