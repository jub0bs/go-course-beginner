package main

import "fmt"

// START OMIT
type Location struct {
	Lat  float64
	Long float64
}

func main() {
	strasbourg := Location{48.5734, 7.7521}       // unkeyed literal: bad!
	paris := Location{Lat: 48.8566, Long: 2.3522} // keyed literal: better
	fmt.Println(strasbourg, paris)
}

// END OMIT
