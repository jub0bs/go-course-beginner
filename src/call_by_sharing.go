package main

import (
	"fmt"
)

// START OMIT
func main() {
	enToFr := map[string]string{}
	upsert(enToFr, "one", "un") // HL
	fmt.Printf("at the end of 'main': %v\n", enToFr)
}

func upsert(m map[string]string, k string, v string) {
	m[k] = v // HL
	fmt.Printf("at the end of 'add': %v\n", m)
}

// END OMIT
