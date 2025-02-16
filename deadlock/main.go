package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 6 threads trying to acquire 2 resources in a circular way. Simulate deadlock.

	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println(id)
		}(i)
	}
	wg.Wait()
}
