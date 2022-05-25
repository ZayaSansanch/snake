package screen

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	fmt.Print("\033[2J")
}

func MoveCursor(x, y int) {
	fmt.Printf("\033[%d;%dH", x, y)
}

func tput(arg string) error {
	cmd := exec.Command("tput", arg)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func Clear() {
	ClearScreen()
	MoveCursor(0, 0)
	tput("civis") // hide cursor
}

func Restore() {
	MoveCursor(0, 0)
	ClearScreen()
	tput("cvvis") // show cursor
}
