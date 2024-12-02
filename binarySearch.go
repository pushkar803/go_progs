package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	target := 3

	x := binarySearch(arr, target)
	fmt.Println(x)
}

func binarySearch(arr []int, target int) int {

	return bsHelper(arr, target, 0, len(arr)-1)

}

func bsHelper(arr []int, target int, left int, right int) int {

	if left > right {
		return -1
	}

	middle := (left + right) / 2
	pmatch := arr[middle]

	if target == pmatch {
		return middle
	} else if target < pmatch {
		return bsHelper(arr, target, left, middle-1)
	}

	return bsHelper(arr, target, middle+1, right)

}
