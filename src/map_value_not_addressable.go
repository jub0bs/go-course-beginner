package main

import (
	"fmt"
)

func main() {
	m := map[string]string{"red": "#ff0000", "blue": "#0000ff", "empty": ""}
	// START OMIT
	v := &m["blue"] // compilation error
	// END OMIT
	fmt.Println(v)
}
