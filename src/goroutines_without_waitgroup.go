package main

import "fmt"

func grindCoffeeBeans() {
	fmt.Println("Grinding coffee beans...")
}

func frothMilk() {
	fmt.Println("Frothing milk...")
}

func main() {
	go grindCoffeeBeans() // HL
	go frothMilk()        // HL
}
