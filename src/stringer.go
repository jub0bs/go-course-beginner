package main

import "fmt"

// START OMIT
type Location struct {
	Lat  float64
	Long float64
}

func (l Location) String() string { // HL
	const tmpl = "%f°N, %f°W"
	return fmt.Sprintf(tmpl, l.Lat, l.Long)
}

func main() {
	paris := Location{Lat: 48.8566, Long: 2.3522}
	fmt.Println(paris)
}

// END OMIT
