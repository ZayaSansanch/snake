package game

import (
	"fmt"
	"time"

	"github.com/ZayaSansanch/snake/screen"
	"github.com/mattn/go-tty"
)

const (
	NoMove    = 0
	MoveUp    = 1
	MoveDown  = 2
	MoveLeft  = 3
	MoveRight = 4
)

func Loop() {
	tty, err := tty.Open()
	if err != nil {
		panic(err)
	}
	defer tty.Close()

	width, height, err := tty.Size()
	if err != nil {
		panic(err)
	}

	var (
		move = NoMove
		stop = false
	)

	go func() {
		for !stop {
			r, err := tty.ReadRune()
			if err != nil {
				panic(err)
			}

			switch r {
			case 13:
				stop = true
			case 119:
				move = MoveUp
			case 115:
				move = MoveDown
			case 97:
				move = MoveLeft
			case 100:
				move = MoveRight
			}
			// fmt.Printf("Key press => '%s' [%d]\n", string(r), r)
		}
	}()

	var (
		x = 10
		y = 10
	)

	for !stop {
		screen.MoveCursor(0, 0)
		fmt.Printf("Direction: %d\n", move)

		screen.MoveCursor(y, x)
		fmt.Print(" ")

		switch move {
		case MoveUp:
			y--
			if y < 0 {
				y = width - 1
			}
		case MoveDown:
			y++
			if y > width-1 {
				y = 0
			}
		case MoveLeft:
			x--
			if x < 0 {
				x = height - 1
			}
		case MoveRight:
			x++
			if x > height-1 {
				x = 0
			}
		}

		screen.MoveCursor(y, x)
		fmt.Print("@")

		time.Sleep(40 * time.Millisecond)
	}
}
