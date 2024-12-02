package main

import "fmt"

func multply(x int) func(y int) int {
	return func(y int) int {
		return x * y
	}
}

func main13() {
	double := multply(2)
	tripple := multply(3)

	dans := double(5)
	tans := tripple(5)

	fmt.Println(dans)
	fmt.Println(tans)

}
