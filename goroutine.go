package main

import (
	"fmt"
	"time"
)

func main4() {

	ch := make(chan int)

	go heavyTask(ch)

	msg := <-ch
	fmt.Println(msg)

}

func heavyTask(ch chan<- int) {

	time.Sleep(2 * time.Second)
	ch <- 3

}
