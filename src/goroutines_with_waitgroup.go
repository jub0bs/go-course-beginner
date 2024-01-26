package main

import (
	"fmt"
	"sync"
)

func grindCoffeeBeans() {
	fmt.Println("Grinding coffee beans...")
}

func frothMilk() {
	fmt.Println("Frothing milk...")
}

// START OMIT
func main() {
	var wg sync.WaitGroup // HL

	wg.Add(1) // HL
	go func() {
		defer wg.Done() // HL
		grindCoffeeBeans()
	}()

	wg.Add(1) // HL
	go func() {
		defer wg.Done() // HL
		frothMilk()
	}()

	wg.Wait() // HL
}

// END OMIT
