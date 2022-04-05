package board

import (
	"math/rand"
	"time"
)

// Size of board
const Size = 9

// Board Difficulty
const (
	Beginner = byte(20)
	Easy     = byte(30)
	Medium   = byte(40)
	Hard     = byte(50)
	VeryHard = byte(60)
)

// Value is a cell value in a sudoku board
type Value int

// Point2 is a position in a 2d array
type Point2 struct{ X, Y int }

// Board is an interface for a sudoku board
type Board interface {
	// Get returns the cell value at the given position
	Get(pos Point2) Value

	// Set sets the cell value at the given position
	Set(pos Point2, value Value)

	// IsPredefined returns if the cell value is defined by system
	IsPredefined(pos Point2) bool

	// IsCorrect returns if the cell value is correct value
	IsCorrect(pos Point2) bool

	// GetConflicts returns positions of cells that has the given value
	// in row, column and subsquare of the given position
	GetConflicts(pos Point2, value Value) map[Point2]struct{}

	// GetPositions returns positions of cells
	// that has the given value in complete board
	GetPositions(value Value) map[Point2]struct{}
}

// New returns a new board instance
func New(difficulty byte) Board {
	firstRow := randomRow()
	complete, _ := Solve(&[Size][Size]Value{firstRow})
	grid := *complete

	for difficulty > 0 {
		pos := randomPos()
		if grid[pos.Y][pos.X] != 0 {
			grid[pos.Y][pos.X] = 0
			difficulty--
		}
	}

	return NewCustom(grid, *complete)
}

// NewCustom returns a new board instance with custom values
func NewCustom(grid [Size][Size]Value, complete [Size][Size]Value) Board {
	board := &board{}

	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			board[i][j] = cell{
				value:      grid[i][j],
				predefined: grid[i][j] != 0,
				correct:    complete[i][j],
			}
		}
	}

	return board
}

type cell struct {
	value      Value
	predefined bool
	correct    Value
}

type board [Size][Size]cell

func (board *board) Get(pos Point2) Value {
	return board[pos.Y][pos.X].value
}

func (board *board) Set(pos Point2, value Value) {
	if !board.IsPredefined(pos) {
		board[pos.Y][pos.X].value = value
	}
}

func (board *board) IsPredefined(pos Point2) bool {
	return board[pos.Y][pos.X].predefined
}

func (board *board) IsCorrect(pos Point2) bool {
	return board.Get(pos) == board[pos.Y][pos.X].correct
}

func (board *board) GetConflicts(pos Point2, value Value) map[Point2]struct{} {
	values := map[Point2]struct{}{}

	// check row
	for i := 0; i < Size; i++ {
		cPos := Point2{i, pos.Y}
		if board.Get(cPos) == value {
			values[cPos] = struct{}{}
		}
	}

	// check column
	for i := 0; i < Size; i++ {
		cPos := Point2{pos.X, i}
		if board.Get(cPos) == value {
			values[cPos] = struct{}{}
		}
	}

	// check subsquare
	sPos := Point2{
		X: int(pos.X/3) * 3,
		Y: int(pos.Y/3) * 3,
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cPos := Point2{sPos.X + j, sPos.Y + i}
			if board.Get(cPos) == value {
				values[cPos] = struct{}{}
			}
		}
	}

	delete(values, pos)
	return values
}

func (board *board) GetPositions(value Value) map[Point2]struct{} {
	values := map[Point2]struct{}{}

	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			pos := Point2{j, i}
			if board.Get(pos) == value {
				values[pos] = struct{}{}
			}
		}
	}

	return values
}

// randomRow returns a randomly shuffled row
func randomRow() [Size]Value {
	row := [Size]Value{1, 2, 3, 4, 5, 6, 7, 8, 9}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(9, func(i, j int) {
		row[i], row[j] = row[j], row[i]
	})

	return row
}

// randomPos returns a random position on the board
func randomPos() Point2 {
	rand.Seed(time.Now().UnixNano())

	return Point2{rand.Intn(Size), rand.Intn(Size)}
}
