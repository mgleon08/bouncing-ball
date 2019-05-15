package main

import "fmt"

func main() {

	const (
		width     = 20
		height    = 12
		cellEmpty = ' '
		cellBall  = '⚾'
	)

	// rune == int32 == ''
	var cell rune

	// create board
	board := make([][]bool, width)
	for columnIndex := range board {
		board[columnIndex] = make([]bool, height)
	}

	board[2][3] = true

	for y := range board[0] {
		for x := range board {
			cell = cellEmpty
			if board[x][y] {
				cell = cellBall
			}
			fmt.Print(string(cell), " ")
		}
		fmt.Println()
	}
}
