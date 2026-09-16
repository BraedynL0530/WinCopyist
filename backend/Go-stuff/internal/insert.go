package internal

// inserts text in detected box

import (
	robot "github.com/go-vgo/robotgo"
)

type InsertionData struct {
	text     string
	length   int
	approved bool // tab pressed or not, possibly useful data/telemetry whatever its called
	relevant bool // thumbs up in gui = yes thumbs down = no, for gui/llm later
}

func findBox() (x, y int) { // needs to find where cursor is and return int/string of what box
	return 1, 0 //temp
}
func Insert(text string, length int) {
	robot.WriteAll(text)
	robot.CmdV()
}
