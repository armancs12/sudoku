package ui

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

// TextWidget is an ui widget for texts
type TextWidget struct {
	String string

	AlignCenter bool
	AlignRight  bool

	Color      tcell.Color
	Background tcell.Color

	lines []string
	width int
}

// Draw draws the text to the terminal
func (tw *TextWidget) Draw(screen tcell.Screen, x, y int) {
	style := tcell.StyleDefault.
		Background(tw.Background).
		Foreground(tw.Color)

	for i, line := range tw.getLines() {
		line = tw.formatString(line)
		for j, char := range []rune(line) {
			if char != ' ' {
				screen.SetContent(x+j, y+i, char, nil, style)
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
