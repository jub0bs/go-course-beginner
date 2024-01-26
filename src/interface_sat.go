package main

import "fmt"

// START OMIT
type Climber interface {
	Climb(meters int) // HL
}

type Mountaineer struct {
	Name     string
	Altitude int
}

func (m *Mountaineer) Climb(meters int) { // HL
	m.Altitude += meters
	fmt.Printf("%s climbs %dm to reach altitude %dm\n", m.Name, meters, m.Altitude)
}

func main() {
	var c Climber = &Mountaineer{Name: "Lynn Hill"} // HL
	fmt.Println(c)
}

// END OMIT
