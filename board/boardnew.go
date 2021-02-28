package board

// Value is a cell value in a sudoku board
type Value int

// Point2 is a position in a 2d array
type Point2 struct{ X, Y int }

// IBoard is an inteface for a sudoku board
type IBoard interface {

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
	GetConflicts(pos Point2, value Value) []Point2

	// GetPositions returns positions of cells
	// that has the given value in complete board
	GetPositions(value Value) []Point2
}

// ===========================================
type board struct{}

// New returns a new board instance
func New() IBoard {
	return nil
}

// NewCustom returns a new board instance with custom values
func NewCustom(grid [9][9]Value, complete [9][9]Value) IBoard {
	return nil
}

func (board *board) Get(pos Point2) Value {
	return 0
}

func (board *board) Set(pos Point2, value Value) {}

func (board *board) IsPredefined(pos Point2) bool {
	return false
}

func (board *board) IsCorrect(pos Point2) bool {
	return false
}

func (board *board) GetConflicts(pos Point2, value Value) []Point2 {
	return nil
}

func (board *board) GetPositions(value Value) []Point2 {
	return nil
}
