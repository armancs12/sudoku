package ui

// BoxWidget is an simple ui wrapper for other widgets
type BoxWidget struct {
	Child Widget

	HAlign    byte
	MinWidth  int
	MinHeight int

	PaddingTop    int
	PaddingBottom int
	PaddingLeft   int
	PaddingRight  int

	Fill  bool
	Color ColorPair
}

const (
	runeULCorner = '┌'
	runeURCorner = '┐'
	runeLLCorner = '└'
	runeLRCorner = '┘'
	runeHLine    = '─'
	runeVLine    = '│'
)

// Draw draws the box widget to the terminal
func (bw *BoxWidget) Draw(context Context, x, y int) {
	width := bw.Width() - 1
	height := bw.Height() - 1

	context.StyleFG(bw.Color.FG)
	context.StyleBG(bw.Color.BG)

	// Draw corners
	context.SetContent(x, y, runeULCorner)
	context.SetContent(x+width, y, runeURCorner)
	context.SetContent(x, y+height, runeLLCorner)
	context.SetContent(x+width, y+height, runeLRCorner)

	// Draw borders
	for i := x + 1; i < width+x; i++ {
		context.SetContent(i, y, runeHLine)
		context.SetContent(i, y+height, runeHLine)
	}
	for j := y + 1; j < height+y; j++ {
		context.SetContent(x, j, runeVLine)
		context.SetContent(x+width, j, runeVLine)
	}

	if bw.Fill {
		for i := x + 1; i < width+x; i++ {
			for j := y + 1; j < height+y; j++ {
				context.SetContent(i, j, ' ')
			}
		}
	}

	childX, childY := bw.getChildPos(x, y)
	bw.Child.Draw(context, childX, childY)
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

func (bw *BoxWidget) getChildPos(x, y int) (int, int) {
	var childX, childY int

	if bw.HAlign == HAlignCenter {
		childX = bw.PaddingLeft + x +
			(bw.Width()-(bw.PaddingLeft+bw.PaddingRight+bw.Child.Width()))/2
		childY = bw.PaddingTop + y +
			(bw.Height()-(bw.PaddingTop+bw.PaddingBottom+bw.Child.Height()))/2
	} else if bw.HAlign == HAlignEnd {
		childX = x + bw.Width() - (bw.Child.Width() + bw.PaddingRight)
		childY = y + bw.Height() - (bw.Child.Height() + bw.PaddingBottom)
	} else {
		childX = bw.PaddingLeft + x
		childY = bw.PaddingTop + y
	}

	return childX, childY
}
