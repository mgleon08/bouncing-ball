package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
	"github.com/mattn/go-runewidth"
)

func main() {
	const (
		cellEmpty = ' '
		cellBall  = '⚾'
		maxFrames = 1200
		speed     = time.Second / 20
		// initial velocities
		ivx, ivy = 2, 12
	)

	var (
		px, py   int        // ball position
		ppx, ppy int        // previous ball position
		vx, vy   = ivx, ivx // velocities
		cell     rune       // rune == int32 == ''
	)

	// you can get the width and height using the screen package easily:
	width, height := screen.Size()
	// width, height, err := terminal.GetSize(int(os.Stdout.Fd()))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// get the rune width of the ball emoji
	ballWidth := runewidth.RuneWidth(cellBall)

	// adjust the width and height
	width /= ballWidth
	height-- // there is a 1 pixel border in my terminal

	bufLen := (width*2 + 1) * height

	// create board
	board := make([][]bool, width)
	for columnIndex := range board {
		board[columnIndex] = make([]bool, height)
	}

	buf := make([]rune, 0, bufLen)

	// clear the screen
	fmt.Print("\033[2J")

	for i := 0; i < maxFrames; i++ {
		// remove the previous ball
		board[px][py] = false

		// move x, y
		px += vx
		py += vy

		// when the ball hits a border reverse its direction
		if px <= 0 || px >= width-ivx {
			vx *= -1
		}
		if py <= 0 || py >= height-ivx {
			vy *= -1
		}

		// put the new ball
		board[px][py], board[ppx][ppy] = true, false

		// save the previous positions
		ppx, ppy = px, py

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
