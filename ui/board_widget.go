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
	Board     board.IBoard
	CursorPos board.Point2

	CursorColor tcell.Color
	NumberColor tcell.Color
	BorderColor tcell.Color
	Background  tcell.Color
}

// Draw draws the board widget to the terminal
func (bw *BoardWidget) Draw(screen tcell.Screen, x, y int) {
	style := tcell.StyleDefault.
		Background(bw.Background).
		Foreground(bw.BorderColor)
	numStyle := tcell.StyleDefault.
		Background(bw.Background).
		Foreground(bw.NumberColor)

	for i, line := range BoardOutline {
		for j, char := range line {
			if number := bw.checkCell(j, i); number != 0 {
				// Workaround for converting integer to rune
				numChar := rune(fmt.Sprintf("%v", number)[0])

				screen.SetContent(x+j, y+i, numChar, nil, numStyle)
			} else {
				screen.SetContent(x+j, y+i, char, nil, style)
			}
		}
	}

	bw.drawCursor(screen, x, y)
}

// Width returns the width of the board widget
func (bw *BoardWidget) Width() int {
	return BoardWidth
}

// Height returns the height of the board widget
func (bw *BoardWidget) Height() int {
	return BoardHeight
}

// checkCell checks if given x, y is corresponds to a cell of the sudoku board.
// If so, returns it. Else it will return 0
func (bw *BoardWidget) checkCell(x, y int) board.Value {
	if bw.isCellPosition(x, y) {
		return bw.getCellValue(x, y)
	}
	return 0
}

func (*BoardWidget) isCellPosition(x, y int) bool {
	return (x-2)%4 == 0 && (y-1)%2 == 0
}

func (bw *BoardWidget) getCellValue(x, y int) board.Value {
	return bw.Board.Get(board.Point2{
		X: (x - 2) / 4,
		Y: (y - 1) / 2,
	})
}

func (bw *BoardWidget) drawCursor(screen tcell.Screen, x, y int) {
	style := tcell.StyleDefault.
		Background(bw.CursorColor).
		Foreground(bw.NumberColor)

	cellX := x + bw.CursorPos.X*4 + 2
	cellY := y + bw.CursorPos.Y*2 + 1

	char, _, _, _ := screen.GetContent(cellX, cellY)

	screen.SetContent(cellX-1, cellY, ' ', nil, style)
	screen.SetContent(cellX, cellY, char, nil, style)
	screen.SetContent(cellX+1, cellY, ' ', nil, style)
}
