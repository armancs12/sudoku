package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/serhatscode/sudoku/board"
)

// BoardOutline for board widget
var BoardOutline = [][]rune{
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

// BoardHeight is height of BoardOutline
var BoardHeight = len(BoardOutline)

// BoardWidth is width of BoardOutline
var BoardWidth = len(BoardOutline[0])

// BoardWidget is an ui widget for sudoku board representation
type BoardWidget struct {
	Board     board.Board
	CursorPos board.Point2

	CursorBG     tcell.Color
	NumberFG     tcell.Color
	BorderFG     tcell.Color
	PredefinedFG tcell.Color
	ConflictFG   tcell.Color
	Background   tcell.Color
}

// Draw draws the board widget to the terminal
func (bw *BoardWidget) Draw(screen tcell.Screen, x, y int) {
	bw.drawBorders(screen, x, y)
	bw.drawCells(screen, x, y)
}

// Width returns the width of the board widget
func (bw *BoardWidget) Width() int {
	return BoardWidth
}

// Height returns the height of the board widget
func (bw *BoardWidget) Height() int {
	return BoardHeight
}

func (bw *BoardWidget) drawBorders(screen tcell.Screen, x, y int) {
	style := tcell.StyleDefault.
		Background(bw.Background).
		Foreground(bw.BorderFG)

	// Horizontal lines
	for i := 0; i < bw.Height(); i += 2 {
		for j := 0; j < bw.Width(); j++ {
			screen.SetContent(x+j, y+i, BoardOutline[i][j], nil, style)
		}
	}

	// Vertical lines
	for i := 1; i < bw.Height(); i += 2 {
		for j := 0; j < bw.Width(); j += 4 {
			screen.SetContent(x+j, y+i, BoardOutline[i][j], nil, style)
		}
	}
}

func (bw *BoardWidget) drawCells(screen tcell.Screen, x, y int) {
	styles := bw.getCellStyles()

	for i := 0; i < board.Size; i++ {
		for j := 0; j < board.Size; j++ {
			pos := board.Point2{X: j, Y: i}
			cx, cy := bw.gridToScreen(pos)
			char := bw.getCellRune(pos)
			style := *styles[i][j]

			screen.SetContent(x+cx, y+cy, ' ', nil, style)
			screen.SetContent(x+cx+1, y+cy, char, nil, style)
			screen.SetContent(x+cx+2, y+cy, ' ', nil, style)
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
	return rune(fmt.Sprintf("%v", bw.Board.Get(pos))[0])
}

func (bw *BoardWidget) getCellStyles() [board.Size][board.Size]*tcell.Style {
	normal := tcell.StyleDefault.Background(bw.Background).Foreground(bw.NumberFG)
	predefined := tcell.StyleDefault.Background(bw.Background).Foreground(bw.PredefinedFG)
	cursor := tcell.StyleDefault.Background(bw.CursorBG).Foreground(bw.NumberFG)
	conflicts := tcell.StyleDefault.Background(bw.Background).Foreground(bw.ConflictFG)
	wrong := tcell.StyleDefault.Background(bw.Background).Foreground(bw.ConflictFG)

	styles := [board.Size][board.Size]*tcell.Style{}

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
