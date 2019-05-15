package main

import (
	"fmt"
)

func main() {
	const (
		width     = 12
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

	buf := make([]rune, width*height)

	board[0][0] = true

	for i := 0; i < 1000; i++ {
		// reuses buffer
		buf = buf[:0]
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}
		fmt.Println(string(buf))
	}
}
