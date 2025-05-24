package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct{ width, height float64 }

func PrintAreas(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Println(s.Area())
	}
}

// START OMIT
func (r *Rectangle) Area() float64 { return r.width * r.height }

func main() {
	PrintAreas(Rectangle{width: 16, height: 9}) // compilation error // HL
}

// END OMIT
