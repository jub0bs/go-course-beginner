package main

import "fmt"

// START OMIT
type Location struct {
	Lat  float64
	Long float64
}

func main() {
	paris := Location{Lat: 48.8566, Long: 2.3522}
	fmt.Println(paris) // {48.8566 2.3522} // HL
}

// END OMIT
