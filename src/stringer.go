package main

import "fmt"

// START OMIT
type Mountaineer struct {
	Name           string
	AltitudeMeters int
}

func (m *Mountaineer) String() string { // HL
	return fmt.Sprintf("%s, currently at altitude %dm\n", m.Name, m.AltitudeMeters)
}

func main() {
	lynn := Mountaineer{
		Name:           "Lynn Hill",
		AltitudeMeters: 1000,
	}
	fmt.Println(&lynn) // HL
}

// END OMIT
