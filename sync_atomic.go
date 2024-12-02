package main

import (
	"fmt"
	"sync/atomic"
)

func main19() {

	var x int32

	atomic.AddInt32(&x, 1)
	fmt.Println(atomic.LoadInt32(&x))

	atomic.StoreInt32(&x, 100)
	fmt.Println(atomic.LoadInt32(&x))

	atomic.CompareAndSwapInt32(&x, 100, 200)
	fmt.Println(atomic.LoadInt32(&x))
}
