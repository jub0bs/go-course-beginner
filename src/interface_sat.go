package main

import "fmt"

// START OMIT
type Climber interface {
	Climb(meters int) error // HL
}

type Mountaineer struct {
	Name     string
	Altitude int
}

func (m *Mountaineer) Climb(meters int) error { // HL
	m.Altitude += meters
	fmt.Printf("%s climbs %dm to reach altitude %dm\n", m.Name, meters, m.Altitude)
	return nil
}

func main() {
	var c Climber = &Mountaineer{Name: "Lynn Hill"} // HL
	fmt.Println(c)
}

// END OMIT
