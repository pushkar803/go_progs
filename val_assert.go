package main

import "fmt"

func main_15() {
	fmt.Println()

	var x interface{} = "d"
	_, ok := x.(string)
	if ok {
		fmt.Println("its string")
	}
}
