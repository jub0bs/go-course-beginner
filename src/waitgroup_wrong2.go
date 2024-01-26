package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go grindCoffeeBeans(wg)
	wg.Wait()
}

func grindCoffeeBeans(wg sync.WaitGroup) { // incorrect: this wg is only a copy main's wg // HL
	defer wg.Done()
	fmt.Println("Grinding coffee beans...")
}

// END OMIT
