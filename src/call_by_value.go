package main

import (
	"fmt"
)

// START OMIT
func main() {
	i := 41
	incr(i)
	fmt.Printf("main: i = %d; &i = %p\n", i, &i)
}

func incr(i int) {
	i++
	fmt.Printf("incr: i = %d; &i = %p\n", i, &i)
}

// END OMIT
