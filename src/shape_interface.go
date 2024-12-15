package main

import (
	"fmt"
	"math"
)

// START1 OMIT
type Shape interface { // HL
	Area() float64 // HL
}

type Rectangle struct{ width, height float64 }

func (r Rectangle) Area() float64 { // enough for Rectangle to implement Shape! // HL
	return r.width * r.height
}

type Circle struct{ radius float64 }

func (c Circle) Area() float64 { // enough for Circle to implement Shape! // HL
	return math.Pi * c.radius * c.radius
}

// END1 OMIT

// START2 OMIT
func PrintAreas(shapes ...Shape) { // HL
	for _, s := range shapes {
		fmt.Println(s.Area())
	}
}

func main() {
	rect := Rectangle{width: 16, height: 9}
	circle := Circle{radius: 10}
	PrintAreas(rect, circle) // HL
}

// END2 OMIT
