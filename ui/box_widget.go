package ui

import (
	"github.com/gdamore/tcell/v2"
)

// BoxWidget is an simple ui wrapper for other widgets
type BoxWidget struct {
	Child Widget

	AlignCenter bool
	MinWidth    int
	MinHeight   int

	PaddingTop    int
	PaddingBottom int
	PaddingLeft   int
	PaddingRight  int

	Fill       bool
	Color      tcell.Color
	Background tcell.Color
}

// Draw draws the box widget to the terminal
func (bw *BoxWidget) Draw(screen tcell.Screen, x, y int) {
	width := bw.Width() - 1
	height := bw.Height() - 1
	style := tcell.StyleDefault.Foreground(bw.Color).Background(bw.Background)

	// Draw corners
	screen.SetContent(x, y, tcell.RuneULCorner, nil, style)
	screen.SetContent(x+width, y, tcell.RuneURCorner, nil, style)
	screen.SetContent(x, y+height, tcell.RuneLLCorner, nil, style)
	screen.SetContent(x+width, y+height, tcell.RuneLRCorner, nil, style)

	// Draw borders
	for i := x + 1; i < width+x; i++ {
		screen.SetContent(i, y, tcell.RuneHLine, nil, style)
		screen.SetContent(i, y+height, tcell.RuneHLine, nil, style)
	}
	for j := y + 1; j < height+y; j++ {
		screen.SetContent(x, j, tcell.RuneVLine, nil, style)
		screen.SetContent(x+width, j, tcell.RuneVLine, nil, style)
	}

	if bw.Fill {
		for i := x + 1; i < width+x; i++ {
			for j := y + 1; j < height+y; j++ {
				screen.SetContent(i, j, ' ', nil, style)
			}
		}
	}

	if bw.AlignCenter {
		xStart := bw.PaddingLeft + x +
			(bw.Width()-(bw.PaddingLeft+bw.PaddingRight+bw.Child.Width()))/2
		yStart := bw.PaddingTop + y +
			(bw.Height()-(bw.PaddingTop+bw.PaddingBottom+bw.Child.Height()))/2

		bw.Child.Draw(screen, xStart, yStart)
	} else {
		bw.Child.Draw(screen, bw.PaddingLeft+x, bw.PaddingTop+y)
	}
}

// Width returns width of the box widget
func (bw *BoxWidget) Width() int {
	width := bw.MinWidth
	if bw.PaddingLeft+bw.PaddingRight+bw.Child.Width() > width {
		width = bw.PaddingLeft + bw.PaddingRight + bw.Child.Width()
	}
	return width
}

// Height returns height of the box widget
func (bw *BoxWidget) Height() int {
	height := bw.MinHeight
	if bw.PaddingTop+bw.PaddingBottom+bw.Child.Height() > height {
		height = bw.PaddingTop + bw.PaddingBottom + bw.Child.Height()
	}
	return height
}
