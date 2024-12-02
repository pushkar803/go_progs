package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	id  int
	val int
}

func main6() {

	numOfTasks := 5
	numOfWorkers := 3

	taskCh := make(chan Task, numOfTasks)
	resultCh := make(chan int, numOfTasks)

	var wg sync.WaitGroup

	for i := 1; i <= numOfWorkers; i++ {
		wg.Add(1)
		go startWorker(i, taskCh, resultCh, &wg)
	}

	for i := 1; i <= numOfTasks; i++ {
		taskCh <- Task{
			id:  i,
			val: i,
		}
	}

	close(taskCh)
	wg.Wait()
	close(resultCh)

	fmt.Println()
	for result := range resultCh {
		fmt.Printf("Received result: %d\n", result)
	}

}

func startWorker(id int, taskCh <-chan Task, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskCh {

		fmt.Printf("worker %d started task %d\n", id, task.id)
		time.Sleep(2 * time.Second)
		res := task.val * 2
		resultCh <- res
		fmt.Printf("worker %d finished task %d\n", id, task.id)

	}
}
