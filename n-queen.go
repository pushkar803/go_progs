package main

import "fmt"

func main16() {

	rows, columns := 4, 4

	var board = make([][]int, rows)

	for i := 0; i < rows; i++ {
		col := make([]int, columns)
		board[i] = col
	}

	for i := 0; i < rows; i++ {

	}

	printBoard(board)
}

func printBoard(board [][]int) {

	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}

}
