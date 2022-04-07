package ui

import (
	"fmt"
	"strings"

	"github.com/serhatsdev/sudoku/game/theme"
)

// TextWidget is an ui widget for texts
type TextWidget struct {
	String string

	AlignCenter bool
	AlignRight  bool

	Color theme.ColorPair

	lines []string
	width int
}

// Draw draws the text to the terminal
func (tw *TextWidget) Draw(context Context, x, y int) {
	context.StyleFG(tw.Color.FG)
	context.StyleBG(tw.Color.BG)

	for i, line := range tw.getLines() {
		line = tw.formatString(line)
		for j, char := range []rune(line) {
			if char != ' ' {
				context.SetContent(x+j, y+i, char)
			}
		}
	}
}

// Width returns the width of the longest text line
func (tw *TextWidget) Width() int {
	if tw.width == 0 {
		for _, line := range tw.getLines() {
			if len(line) > tw.width {
				tw.width = len(line)
			}
		}
	}
	return tw.width
}

// Height returns number of lines
func (tw *TextWidget) Height() int {
	return len(tw.getLines())
}

func (tw *TextWidget) getLines() []string {
	if tw.lines == nil {
		tw.lines = strings.Split(tw.String, "\n")
	}
	return tw.lines
}

func (tw *TextWidget) formatString(str string) string {
	width := tw.Width()
	if tw.AlignCenter {
		length := len(str)
		pad := int((length - width) / 2)
		return fmt.Sprintf("%*v%v%-*v", pad, "", str, pad, "")
	} else if tw.AlignRight {
		return fmt.Sprintf("%*v", width, str)
	}
	return fmt.Sprintf("%-*v", width, str)
}
