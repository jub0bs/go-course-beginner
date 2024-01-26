package main

import "fmt"

// START OMIT
func main() {
	ch := make(chan string)

	in := []string{"foo", "bar", "baz"}
	for _, name := range in {
		go func() { ch <- name }()
	}

	var out []string
	for range len(in) { // HL
		out = append(out, <-ch) // HL
	}
	fmt.Println(out)
}

// END OMIT
