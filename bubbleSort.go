package main

import "fmt"

func main3() {
	x := []int{1, 3, 2, 8, 5}
	y := bubbleSort(x)
	fmt.Println(y)
}

func bubbleSort(x []int) []int {
	n := len(x)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
	return x
}
