package sudoku

import "errors"

// SolveBoard solves the board using backtracking algorithm
func SolveBoard(board Board) (*Board, error) {
	column, row, found := board.GetNextEmptyCell()
	if !found {
		return &board, nil
	}

	for i := 1; i < 10; i++ {
		if board.IsGuessValid(column, row, i) {
			board[row][column] = i
			solved, err := SolveBoard(board)
			if err == nil {
				return solved, nil
			}
		}
		board[row][column] = 0
	}

	return nil, errors.New("board is unsolvable")
}
