package main

// START OMIT
import (
	"fmt"
	"time" // HL
)

func main() {
	go grindCoffeeBeans()
	go frothMilk()
	time.Sleep(100 * time.Millisecond) // HL
}

// omitted declarations of functions grindCoffeeBeans and frothMilk
// END OMIT

func grindCoffeeBeans() {
	fmt.Println("Grinding coffee beans...")
}

func frothMilk() {
	fmt.Println("Frothing milk...")
}
