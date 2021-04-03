package ui

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
	Color      string
	Background string
}

const (
	RuneULCorner = '┌'
	RuneURCorner = '┐'
	RuneLLCorner = '└'
	RuneLRCorner = '┘'
	RuneHLine    = '─'
	RuneVLine    = '│'
)

// Draw draws the box widget to the terminal
func (bw *BoxWidget) Draw(context Context, x, y int) {
	width := bw.Width() - 1
	height := bw.Height() - 1

	context.StyleFG(bw.Color)
	context.StyleBG(bw.Background)

	// Draw corners
	context.SetContent(x, y, RuneULCorner)
	context.SetContent(x+width, y, RuneURCorner)
	context.SetContent(x, y+height, RuneLLCorner)
	context.SetContent(x+width, y+height, RuneLRCorner)

	// Draw borders
	for i := x + 1; i < width+x; i++ {
		context.SetContent(i, y, RuneHLine)
		context.SetContent(i, y+height, RuneHLine)
	}
	for j := y + 1; j < height+y; j++ {
		context.SetContent(x, j, RuneVLine)
		context.SetContent(x+width, j, RuneVLine)
	}

	if bw.Fill {
		for i := x + 1; i < width+x; i++ {
			for j := y + 1; j < height+y; j++ {
				context.SetContent(i, j, ' ')
			}
		}
	}

	if bw.AlignCenter {
		xStart := bw.PaddingLeft + x +
			(bw.Width()-(bw.PaddingLeft+bw.PaddingRight+bw.Child.Width()))/2
		yStart := bw.PaddingTop + y +
			(bw.Height()-(bw.PaddingTop+bw.PaddingBottom+bw.Child.Height()))/2

		bw.Child.Draw(context, xStart, yStart)
	} else {
		bw.Child.Draw(context, bw.PaddingLeft+x, bw.PaddingTop+y)
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
