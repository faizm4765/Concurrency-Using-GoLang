package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var wgD sync.WaitGroup

type ConcurrentQueue struct {
	queue []int

	// Since mutex needs to be shared by all th eoperations, mutex variable needs to be common, so we need to described at struct level
	mu sync.Mutex
}

func (cq *ConcurrentQueue) Enqueue(item int) {
	cq.mu.Lock()
	cq.queue = append(cq.queue, item)
	cq.mu.Unlock()
	wg.Done()
}

func (cq *ConcurrentQueue) Dequeue() int {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	if len(cq.queue) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}

	item := cq.queue[0]
	cq.queue = cq.queue[1:]
	return item
}

func (cq *ConcurrentQueue) size() int {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	return len(cq.queue)
}

func main() {
	cq := ConcurrentQueue{}

	// create multiple threads inserting values in queue
	for i := 1; i <= 1000000; i++ {
		wg.Add(1)
		go cq.Enqueue(i)
	}

	for i := 1; i <= 1000000; i++ {
		wgD.Add(1)
		go func() {
			cq.Dequeue()
			wgD.Done() // wait group can be decremented when the operation is done. Earlier above wgD.Done() was inside the enqueue method, which is also fine as it was executed post execution of goroutine. This is same case here, only thing is we leverage anonymous method.
		}()
	}

	wg.Wait()
	wgD.Wait()
	fmt.Println(cq.size())
}
