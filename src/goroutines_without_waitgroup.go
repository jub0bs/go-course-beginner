package main

// START OMIT
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

// END OMIT
