package main

import "fmt"

func main14() {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	go func() {
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

}
