package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

// MenuWidget is an ui widget for menu representations
type MenuWidget struct {
	Options     []string
	CursorIndex int
	MinWidth    int

	AlignCenter bool
	AlignRight  bool

	Color            tcell.Color
	Background       tcell.Color
	CursorColor      tcell.Color
	CursorBackground tcell.Color

	width int
}

// Draw draws the menu widget to the terminal
func (mw *MenuWidget) Draw(screen tcell.Screen, x, y int) {
	for i, option := range mw.Options {
		style := mw.getStyleForOption(i)
		option = mw.formatOption(option)

		for j, char := range []rune(option) {
			screen.SetContent(x+j, y+i, char, nil, style)
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

func (mw *MenuWidget) getStyleForOption(index int) tcell.Style {
	if mw.CursorIndex == index {
		return tcell.StyleDefault.
			Background(mw.CursorBackground).
			Foreground(mw.CursorColor)
	}
	return tcell.StyleDefault.
		Background(mw.Background).
		Foreground(mw.Color)
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
