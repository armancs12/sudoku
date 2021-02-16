package ui

import (
	"github.com/gdamore/tcell/v2"
)

// scr to make ui to work
var scr tcell.Screen

// Init sets & initializes screen
func Init(screen tcell.Screen) error {
	err := screen.Init()
	if err != nil {
		return err
	}

	scr = screen
	return nil
}

// Draw draws the given widget to terminal
func Draw(x, y int, widget Widget) {
	widget.Draw(scr, x, y)
	scr.Show()
}

// DrawCenter draws the given widget in center of the terminal
func DrawCenter(widget Widget) {
	width, height := scr.Size()
	x := (width - widget.Width()) / 2
	y := (height - widget.Height()) / 2

	Draw(x, y, widget)
}

// DrawUpLeftCorner draws the given widget to terminal at up left corner
func DrawUpLeftCorner(widget Widget) {
	Draw(0, 0, widget)
}

// DrawUpRightCorner draws the given widget to terminal at up right corner
func DrawUpRightCorner(widget Widget) {
	width, _ := scr.Size()
	x := width - widget.Width()

	Draw(x, 0, widget)
}

// DrawDownLeftCorner draws the given widget to terminal at down left corner
func DrawDownLeftCorner(widget Widget) {
	_, height := scr.Size()
	y := height - widget.Height()

	Draw(0, y, widget)
}

// DrawDownRightCorner draws the given widget to terminal at down right corner
func DrawDownRightCorner(widget Widget) {
	width, height := scr.Size()
	x := width - widget.Width()
	y := height - widget.Height()

	Draw(x, y, widget)
}

// Widget is representation for an ui item
type Widget interface {
	Draw(screen tcell.Screen, x, y int)
	Width() int
	Height() int
}

// NullWidget is an nil replacement for widget interface
type NullWidget struct{}

// Draw does nothing
func (*NullWidget) Draw(screen tcell.Screen, x, y int) {}

// Width returns 0
func (*NullWidget) Width() int { return 0 }

// Height returns 0
func (*NullWidget) Height() int { return 0 }
