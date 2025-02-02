package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var wg sync.WaitGroup
var count int

func incCount() {
	mu.Lock()
	count++ // This is a critical section. We want this part to be acquired by only one goroutine at a time through a lock.
	mu.Unlock()
	wg.Done() // Ensure wg.Done() is executed when the goroutine completes its task. Meaning always call it inside the goroutine.
}

func doCount() {
	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go incCount()
	}
}

func main() {
	count = 0
	doCount()
	wg.Wait()
	fmt.Println(count)
}
