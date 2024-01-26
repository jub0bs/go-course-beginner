package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// START OMIT
	m := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
	}
	for _, k := range slices.Sorted(maps.Keys(m)) {
		fmt.Printf("%s: %q\n", k, m[k])
	}
	// END OMIT
}
