package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu [6]sync.Mutex

func main() {
	// 6 threads trying to acquire 2 resources in a circular way. Simulate deadlock.

	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println(id)
			resource1 := id
			resource2 := (id + 1) % 6

			fmt.Printf("Thread %d trying to acquire resource %d\n", id, resource1)
			mu[resource1].Lock()
			fmt.Printf("Thread %d acquired resource %d\n", id, resource1)

			// Simulate some work
			time.Sleep(1 * time.Second)

			fmt.Printf("Thread %d trying to acquire resource %d\n", id, resource2)
			mu[resource2].Lock()
			fmt.Printf("Thread %d acquired resource %d\n", id, resource2)

			mu[resource1].Unlock()
			mu[resource2].Unlock()
		}(i)
	}

	wg.Wait()
}
