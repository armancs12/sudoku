package ui

const (
	HAlignStart = byte(iota)
	HAlignCenter
	HAlignEnd
	VAlignStart
	VAlignCenter
	VAlignEnd
)

type ColorPair struct {
	FG string
	BG string
}

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
	DrawCenter(widget Widget)
	// Draw draws the given widget at the up right corner of the screen
	DrawAligned(widget Widget, alignments ...byte)

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

// Widget is representation for an ui item
type Widget interface {
	Draw(context Context, x, y int)
	Width() int
	Height() int
}

// NullWidget is an nil replacement for widget interface
type NullWidget struct{}

// Draw does nothing
func (*NullWidget) Draw(context Context, x, y int) {}

// Width returns 0
func (*NullWidget) Width() int { return 0 }

// Height returns 0
func (*NullWidget) Height() int { return 0 }
