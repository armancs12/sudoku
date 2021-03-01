package board

import (
	"errors"
)

// Solve solves the given sudoku grid using backtracking algorithm
func Solve(grid *[Size][Size]Value) (*[Size][Size]Value, error) {
	pos, exist := findNextEmptyCell(grid)
	if !exist {
		return grid, nil
	}

	for i := Value(1); i < 10; i++ {
		if isValid(grid, pos, i) {
			grid[pos.Y][pos.X] = i

			if solved, err := Solve(grid); err == nil {
				return solved, nil
			}
		}
		grid[pos.Y][pos.X] = 0
	}

	return nil, errors.New("board is unsolvable")
}

// findNextEmptyCell returns next empty cell if exist
func findNextEmptyCell(grid *[Size][Size]Value) (Point2, bool) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				return Point2{j, i}, true
			}
		}
	}
	return Point2{0, 0}, false
}

func isValid(grid *[Size][Size]Value, pos Point2, value Value) bool {
	// check row
	for i := 0; i < Size; i++ {
		if grid[pos.Y][i] == value {
			return false
		}
	}

	// check column
	for i := 0; i < Size; i++ {
		if grid[i][pos.X] == value {
			return false
		}
	}

	// check subsquare
	sPos := Point2{
		X: int(pos.X/3) * 3,
		Y: int(pos.Y/3) * 3,
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[sPos.Y+i][sPos.X+j] == value {
				return false
			}
		}
	}

	return true
}
