package main

import "sync"

// START OMIT
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

// END OMIT
