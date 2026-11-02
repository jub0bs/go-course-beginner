package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var wg sync.WaitGroup
	wg.Go(func() { count++ })
	wg.Go(func() { count++ })
	wg.Wait()
	fmt.Println(count)
}
