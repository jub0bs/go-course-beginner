package main

import "fmt"

// START OMIT
func main() {
	fmt.Println(Min(42, 999, -1, 23), Max(44, 99))
}

func Min(z int, nz ...int) int {
	return Reduce(min, z, ns...) // compilation error: min (built-in) must be called // HL
}

func Max(z int, nz ...int) int {
	return Reduce(max, z, ns...) // compilation error: max (built-in) must be called // HL
}

func Reduce(f func(acc, n int) int, z int, ns ...int) int {
	// implementation omitted
	// END OMIT
	for _, n := range ns {
		z = f(z, n)
	}
	return z
}
