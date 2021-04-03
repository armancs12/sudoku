package ui

import (
	"fmt"
)

// MenuWidget is an ui widget for menu representations
type MenuWidget struct {
	Options     []string
	CursorIndex int
	MinWidth    int

	AlignCenter bool
	AlignRight  bool

	Color            string
	Background       string
	CursorColor      string
	CursorBackground string

	width int
}

// Draw draws the menu widget to the terminal
func (mw *MenuWidget) Draw(context Context, x, y int) {
	for i, option := range mw.Options {
		fg, bg := mw.getStyleForOption(i)
		option = mw.formatOption(option)

		context.StyleFG(fg)
		context.StyleBG(bg)
		for j, char := range []rune(option) {
			context.SetContent(x+j, y+i, char)
		}
	}
}

// Width returns the width of the menu widget
func (mw *MenuWidget) Width() int {
	if mw.width == 0 {
		mw.width = mw.MinWidth
		for _, option := range mw.Options {
			if len(option) > mw.width {
				mw.width = len(option)
			}
		}
	}

	return mw.width
}

// Height returns the height of the menu widget
func (mw *MenuWidget) Height() int {
	return len(mw.Options)
}

func (mw *MenuWidget) getStyleForOption(index int) (string, string) {
	if mw.CursorIndex == index {
		return mw.CursorColor, mw.CursorBackground
	}
	return mw.Color, mw.Background
}

func (mw *MenuWidget) formatOption(option string) string {
	width := mw.Width()
	if mw.AlignCenter {
		length := len(option)
		pad := int((length - width) / 2)
		return fmt.Sprintf("%*v%v%-*v", pad, "", option, pad, "")
	} else if mw.AlignRight {
		return fmt.Sprintf("%*v", width, option)
	}
	return fmt.Sprintf("%-*v", width, option)
}
