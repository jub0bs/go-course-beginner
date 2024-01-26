package main

import "fmt"

// START OMIT
// "any" is the type constraint on type parameter E: it means that E can be any arbitrary type
type Node[E any] struct { // HL
	Elem  E        // HL
	Left  *Node[E] // HL
	Right *Node[E] // HL
}

func Size[E any](t *Node[E]) int { // HL
	if t == nil {
		return 0
	}
	return 1 + Size(t.Left) + Size(t.Right)
}

func main() {
	tree := &Node[int]{Elem: 3, Left: &Node[int]{Elem: 13}}
	fmt.Println(Size(tree))
}

// END OMIT
