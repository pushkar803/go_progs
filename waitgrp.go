package main

import (
	"fmt"
	"sync"
)

func main5() {

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go heavyTaskd(i, &wg)
	}
	//time.Sleep(3 * time.Second)
	wg.Wait()
}

func heavyTaskd(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
}
