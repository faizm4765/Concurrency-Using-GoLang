package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value   int64
	version int64
}

var wg sync.WaitGroup

func (c *Counter) Update(newValue int64) {
	defer wg.Done()
	curVal := atomic.LoadInt64(&c.value)

	if atomic.CompareAndSwapInt64(&c.value, curVal, newValue) {
		atomic.AddInt64(&c.version, 1)
		atomic.StoreInt64(&c.value, newValue)
		fmt.Println("Updated value of counter:", c.value)
		return
	}

	fmt.Println("Failed to update counter!")
	return
}

func main() {
	counter := Counter{value: 0, version: 0}
	wg.Add(2)
	go counter.Update(10)
	go counter.Update(20)
	// var i int64
	// for i = 1; i < 10; i++ {
	// 	wg.Add(1)
	// 	go counter.Update(i)
	// }

	wg.Wait()
	fmt.Println("Final value of counter:", counter.value)
}
