package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		count++
	}()
	go func() {
		defer wg.Done()
		count++
	}()
	wg.Wait()
	fmt.Println(count)
}
