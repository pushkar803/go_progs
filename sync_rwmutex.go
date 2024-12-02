package main

import (
	"fmt"
	"sync"
)

type CounterRW struct {
	mu  sync.RWMutex
	val int
}

func (c *CounterRW) read() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.val
}

func (c *CounterRW) write() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func main18() {
	var counter CounterRW

	for range [10]int{} {
		//fmt.Println(i)
		go backgroundWrite(&counter)
		backgroundRead(&counter)
	}

}

func backgroundWrite(counter *CounterRW) {
	counter.write()
}

func backgroundRead(counter *CounterRW) {
	x := counter.read()
	fmt.Println(x)
}
