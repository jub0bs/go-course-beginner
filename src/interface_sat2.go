package main

import "fmt"

type Climber interface {
	Climb(meters int) // HL
}

type Mountaineer struct {
	Name     string
	Altitude int
}

// START OMIT
func (m *Mountaineer) Climb(meters int) { /* ... */ }

func main() {
	var c Climber = Mountaineer{Name: "Lynn Hill"} // compilation error // HL
	fmt.Println(c)
}

// END OMIT
