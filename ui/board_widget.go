package ui

import (
	"github.com/serhatscode/sudoku/board"
)

// BoardHeight is height of BoardOutline
var BoardHeight = len(boardOutline)

// BoardWidth is width of BoardOutline
var BoardWidth = len(boardOutline[0])

// boardOutline for board widget
var boardOutline = [][]rune{
	[]rune("┏━━━┯━━━┯━━━┳━━━┯━━━┯━━━┳━━━┯━━━┯━━━┓"),
	[]rune("┃   │   │   ┃   │   │   ┃   │   │   ┃"),
	[]rune("┠───┼───┼───╂───┼───┼───╂───┼───┼───┨"),
	[]rune("┃   │   │   ┃   │   │   ┃   │   │   ┃"),
	[]rune("┠───┼───┼───╂───┼───┼───╂───┼───┼───┨"),
	[]rune("┃   │   │   ┃   │   │   ┃   │   │   ┃"),
	[]rune("┣━━━┿━━━┿━━━╋━━━┿━━━┿━━━╋━━━┿━━━┿━━━┫"),
	[]rune("┃   │   │   ┃   │   │   ┃   │   │   ┃"),
	[]rune("┠───┼───┼───╂───┼───┼───╂───┼───┼───┨"),
	[]rune("┃   │   │   ┃   │   │   ┃   │   │   ┃"),
	[]rune("┠───┼───┼───╂───┼───┼───╂───┼───┼───┨"),
	[]rune("┃   │   │   ┃   │   │   ┃   │   │   ┃"),
	[]rune("┣━━━┿━━━┿━━━╋━━━┿━━━┿━━━╋━━━┿━━━┿━━━┫"),
	[]rune("┃   │   │   ┃   │   │   ┃   │   │   ┃"),
	[]rune("┠───┼───┼───╂───┼───┼───╂───┼───┼───┨"),
	[]rune("┃   │   │   ┃   │   │   ┃   │   │   ┃"),
	[]rune("┠───┼───┼───╂───┼───┼───╂───┼───┼───┨"),
	[]rune("┃   │   │   ┃   │   │   ┃   │   │   ┃"),
	[]rune("┗━━━┷━━━┷━━━┻━━━┷━━━┷━━━┻━━━┷━━━┷━━━┛"),
}

type boardColors struct {
	fg string
	bg string
}

// BoardWidget is an ui widget for sudoku board representation
type BoardWidget struct {
	Board     board.Board
	CursorPos board.Point2

	CursorBG     string
	NumberFG     string
	BorderFG     string
	PredefinedFG string
	ConflictFG   string
	Background   string
}

// Draw draws the board widget to the terminal
func (bw *BoardWidget) Draw(context Context, x, y int) {
	bw.drawBorders(context, x, y)
	bw.drawCells(context, x, y)
}

// Width returns the width of the board widget
func (bw *BoardWidget) Width() int {
	return BoardWidth
}

// Height returns the height of the board widget
func (bw *BoardWidget) Height() int {
	return BoardHeight
}

func (bw *BoardWidget) drawBorders(context Context, x, y int) {
	context.StyleFG(bw.BorderFG)
	context.StyleBG(bw.Background)

	// Horizontal lines
	for i := 0; i < bw.Height(); i += 2 {
		for j := 0; j < bw.Width(); j++ {
			context.SetContent(x+j, y+i, boardOutline[i][j])
		}
	}

	// Vertical lines
	for i := 1; i < bw.Height(); i += 2 {
		for j := 0; j < bw.Width(); j += 4 {
			context.SetContent(x+j, y+i, boardOutline[i][j])
		}
	}
}

func (bw *BoardWidget) drawCells(context Context, x, y int) {
	styles := bw.getCellStyles()

	for i := 0; i < board.Size; i++ {
		for j := 0; j < board.Size; j++ {
			pos := board.Point2{X: j, Y: i}
			cx, cy := bw.gridToScreen(pos)
			char := bw.getCellRune(pos)
			style := *styles[i][j]

			context.StyleFG(style.fg)
			context.StyleBG(style.bg)

			context.SetContent(x+cx, y+cy, ' ')
			context.SetContent(x+cx+1, y+cy, char)
			context.SetContent(x+cx+2, y+cy, ' ')
		}
	}
}

func (bw *BoardWidget) gridToScreen(pos board.Point2) (int, int) {
	return pos.X*4 + 1, pos.Y*2 + 1
}

func (bw *BoardWidget) getCellRune(pos board.Point2) rune {
	if bw.Board.Get(pos) == 0 {
		return ' '
	}

	return '0' + rune(bw.Board.Get(pos))
}

func (bw *BoardWidget) getCellStyles() [board.Size][board.Size]*boardColors {
	normal := boardColors{bw.NumberFG, bw.Background}
	predefined := boardColors{bw.PredefinedFG, bw.Background}
	cursor := boardColors{bw.NumberFG, bw.CursorBG}
	conflicts := boardColors{bw.ConflictFG, bw.Background}
	wrong := boardColors{bw.ConflictFG, bw.Background}

	styles := [board.Size][board.Size]*boardColors{}

	styles[bw.CursorPos.Y][bw.CursorPos.X] = &cursor

	// Set conflict style
	value := bw.Board.Get(bw.CursorPos)
	if value != 0 {
		for conflict := range bw.Board.GetConflicts(bw.CursorPos, value) {
			styles[conflict.Y][conflict.X] = &conflicts
		}
	}

	// Set other styles
	for i := 0; i < board.Size; i++ {
		for j := 0; j < board.Size; j++ {
			if styles[i][j] != nil {
				continue
			}

			pos := board.Point2{X: j, Y: i}
			if bw.Board.IsPredefined(pos) {
				styles[pos.Y][pos.X] = &predefined
			} else if bw.Board.Get(pos) != 0 && !bw.Board.IsCorrect(pos) {
				styles[pos.Y][pos.X] = &wrong
			} else {
				styles[pos.Y][pos.X] = &normal
			}
		}
	}
	return styles
}
