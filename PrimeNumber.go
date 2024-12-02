package main

import (
	"fmt"
	"math"
)

func main2() {
	n := 3
	x := checkIfPrimeNumber(n)
	fmt.Println(x)
}

func checkIfPrimeNumber(n int) bool {

	if n < 2 {
		return false
	}

	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
