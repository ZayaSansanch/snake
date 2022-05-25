package main

import (
	"github.com/ZayaSansanch/snake/game"
	"github.com/ZayaSansanch/snake/screen"
)

func main() {
	screen.Clear()
	game.Loop()
	screen.Restore()
}
