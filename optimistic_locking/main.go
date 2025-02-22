package main

import "fmt"

type Counter struct {
	value   int
	version int
}

func (c *Counter) Update(newValue int) {
	c.value = newValue
	c.version++
	fmt.Println("Updated value of counter:", c.value)
}

func main() {
	counter := Counter{value: 0, version: 0}
	counter.Update(10)
	counter.Update(20)

	fmt.Println("Final value of counter:", counter.value)
}
