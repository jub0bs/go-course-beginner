package main

import "fmt"

// START OMIT
// "any" is the type constraint on type parameter E: it means that E can be any arbitrary type
type Tree[E any] struct { // HL
	Left  *Tree[E] // HL
	Elem  E        // HL
	Right *Tree[E] // HL
}

func Size[E any](t *Tree[E]) int { // HL
	if t == nil {
		return 0
	}
	return 1 + Size(t.Left) + Size(t.Right)
}

func main() {
	tree := &Tree[int]{Elem: 3, Left: &Tree[int]{Elem: 13}}
	fmt.Println(Size(tree))
}

// END OMIT
