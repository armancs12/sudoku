package sudoku

// Board is representing sudoku board
type Board [9][9]int

// ColumnHasNumber checks if the column has same number
func (board *Board) ColumnHasNumber(column, number int) bool {
	for _, row := range board {
		if number == row[column] {
			return true
		}
	}
	return false
}

// RowHasNumber checks if the row has same number
func (board *Board) RowHasNumber(row, number int) bool {
	for _, num := range board[row] {
		if number == num {
			return true
		}
	}
	return false
}

// BoxHasNumber checks if the box has same number
func (board *Board) BoxHasNumber(boxColumn, boxRow, number int) bool {
	boxSlice := getBoxSlice(*board, boxColumn, boxRow)
	for _, row := range boxSlice {
		for _, num := range row {
			if number == num {
				return true
			}
		}
	}
	return false
}

// IsGuessValid checks if guess not violates board rules
func (board *Board) IsGuessValid(column, row int, guess int) bool {
	boxColumn, boxRow := getBoxCoordinate(column, row)
	return !(board.BoxHasNumber(boxColumn, boxRow, guess) ||
		board.ColumnHasNumber(column, guess) ||
		board.RowHasNumber(row, guess))
}

// GetNextEmptyCell gets first empty cell if exist
func (board *Board) GetNextEmptyCell() (int, int, bool) {
	for i, row := range board {
		for j, num := range row {
			if num == 0 {
				return j, i, true
			}
		}
	}
	return 0, 0, false
}

func getBoxSlice(board Board, boxColumn, boxRow int) [][]int {
	slice := [][]int{}
	for i := 0; i < 3; i++ {
		slice = append(slice, board[(boxRow*3 + i)][(boxColumn*3):(boxColumn*3+3)])
	}

	return slice
}

func getBoxCoordinate(column, row int) (int, int) {
	return (column / 3), (row / 3)
}
