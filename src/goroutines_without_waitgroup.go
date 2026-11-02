package main

import "fmt"

func grindCoffeeBeans(grams float64) {
	fmt.Printf("Grinding %f grams of coffee beans...\n", grams)
}

func frothMilk(ml float64) {
	fmt.Printf("Frothing %f ml of milk...\n", ml)
}

func main() {
	go grindCoffeeBeans(7) // HL
	go frothMilk(30)       // HL
}
