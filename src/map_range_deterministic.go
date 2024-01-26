package main

import (
	"fmt"
	"sort"

	"golang.org/x/exp/maps"
)

func main() {
	// START OMIT
	m := map[string]string{"red": "#ff0000", "blue": "#0000ff", "green": "#00ff00"}
	keys := maps.Keys(m)
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: %q\n", k, m[k])
	}
	// END OMIT
}
