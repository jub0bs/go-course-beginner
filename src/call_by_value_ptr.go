package main

import (
	"fmt"
)

// START OMIT
func main() {
	i := 41
	incr(&i)
	fmt.Printf("main:  i = %d; &i = %p\n", i, &i)
}

func incr(p *int) {
	*p++
	fmt.Printf("incr: *p = %d;  p = %p\n", *p, p)
}

// END OMIT
