package main

import (
	"fmt"
	"time"
)

func main7() {
	fmt.Println("main program started")
	ch := make(chan string)
	go task1(ch)
	//msg := <-ch
	for msg := range ch {
		fmt.Println(msg)
	}
	fmt.Println("main program ended")

}

func task1(chs chan<- string) {
	fmt.Println("task started")
	time.Sleep(2 * time.Second)
	chs <- "hello 1"
	chs <- "hello 2"
	chs <- "hello 3"
	fmt.Println("task ended")
	close(chs)
}
