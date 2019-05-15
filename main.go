package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		width     = 50
		height    = 12
		cellEmpty = ' '
		cellBall  = '⚾'
	)

	var (
		px, py    int
		vx, vy    = 1, 1 // velocities
		cell      rune   // rune == int32 == ''
		maxFrames = 1200
		speed     = time.Second / 20
	)

	// create board
	board := make([][]bool, width)
	for columnIndex := range board {
		board[columnIndex] = make([]bool, height)
	}

	buf := make([]rune, width*height)

	// clear the screen
	fmt.Print("\033[2J")

	for i := 0; i < maxFrames; i++ {
		// remove the previous ball
		board[px][py] = false

		// move x, y
		px += vx
		py += vy

		// when the ball hits a border reverse its direction
		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		// put the new ball
		board[px][py] = true

		// reuses buffer
		buf = buf[:0]

		// draw the board
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

		// move top left
		fmt.Print("\033[H")
		fmt.Println(string(buf))

		// slow down the animation
		time.Sleep(speed)
	}
}
