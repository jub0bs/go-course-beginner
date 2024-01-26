package main

import (
	"sync"
)

func main() {
	// START OMIT
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
	// END OMIT
}
