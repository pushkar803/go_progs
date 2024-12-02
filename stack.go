package main

import "fmt"

type StackManager struct {
	stack []int
}

func (sm *StackManager) push(data int) {
	sm.stack = append(sm.stack, data)
}

func (sm *StackManager) pop() {
	sm.stack = sm.stack[:len(sm.stack)-1]
}

func (sm *StackManager) print() {
	fmt.Println(sm.stack)
}

func main21() {

	var sm StackManager

	for i := 0; i < 10; i++ {
		sm.push(i)
	}
	for i := 0; i < 3; i++ {
		sm.pop()
	}

	sm.print()
}
