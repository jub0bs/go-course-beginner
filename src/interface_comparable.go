package main

import "fmt"

func main() {
	// START OMIT
	var i, j any = 42, 42
	fmt.Println(i == j) // true: same dynamic type and value
	j = 3.14
	fmt.Println(i == j) // false: different dynamic types
	i, j = 42, 99
	fmt.Println(i == j) // false: same dynamic type but different dynamic values
	i, j = []int{42}, []int{99}
	fmt.Println(i == j) // panic: runtime error: comparing uncomparable type []int
	// END OMIT
}
