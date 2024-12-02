package main

import (
	"fmt"
	"time"
)

func main12() {
	ch := make(chan int) // unbuffered channel

	// This will result in deadlock because both goroutines are blocked
	go func() {
		ch <- 42 // blocks until main goroutine receives
	}()
	time.Sleep(2 * time.Second)
	fmt.Println(<-ch) // blocks until a value is sent to the channel
}
