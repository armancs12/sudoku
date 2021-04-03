package ui

import (
	"github.com/gdamore/tcell/v2"
)

type Client interface {
	// Start starts the client
	Start() error
	// Stop stops the client
	Stop()

	// Size returns width and height of client screen
	Size() (int, int)

	// OnResize takes a function to run when the screen resized
	OnResize(func(width, height int))
	// OnKeyPress takes a function to run when a key pressed
	OnKeyPress(func(key string))

	// Draw draws the given widget at the given position
	Draw(x, y int, widget Widget)
	// Draw draws the given widget at the center of the screen
	DrawCenter(Widget)
	// Draw draws the given widget at the up right corner of the screen
	DrawUpRightCorner(Widget)

	// Context returns context to manipulate the screen
	Context() Context
}

type Context interface {
	// StyleFG sets the current foreground color
	StyleFG(color string)
	// StyleBG sets the current background color
	StyleBG(color string)

	// SetContent draws the given char to the given position
	// with the current style
	SetContent(x, y int, char rune)

	// Show makes all content changes visible
	Show()

	// Clear clears the screen
	Clear()
}

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
