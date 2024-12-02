package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu  sync.Mutex
	val int
}

func (c *Counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func (c *Counter) getVal() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}

func main17() {

	var counter Counter

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go backgroundInc(&counter, &wg)

	}

	wg.Wait()

	x := counter.getVal()
	fmt.Println(x)

}

func backgroundInc(counter *Counter, wg *sync.WaitGroup) {
	wg.Done()
	counter.increment()
}
