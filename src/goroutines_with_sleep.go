// START OMIT
package main

import (
	"fmt"
	"time" // HL
)

func main() {
	go grindCoffeeBeans(7)
	go frothMilk(30)
	time.Sleep(100 * time.Millisecond) // HL
}

// omitted declarations of functions grindCoffeeBeans and frothMilk
// END OMIT

func grindCoffeeBeans(grams float64) {
	fmt.Printf("Grinding %f grams of coffee beans...\n", grams)
}

func frothMilk(ml float64) {
	fmt.Printf("Frothing %f ml of milk...\n", ml)
}
