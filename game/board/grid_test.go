package board_test

import (
	"testing"

	"github.com/serhatsdev/sudoku/game/board"
)

func TestGridGetRow(t *testing.T) {
	grid := board.Grid{
		{6, 2, 4, 7, 9, 3, 5, 8, 1},
		{7, 5, 1, 8, 4, 6, 9, 3, 2},
		{8, 3, 9, 1, 2, 5, 6, 4, 7},
		{4, 8, 3, 9, 5, 7, 2, 1, 6},
		{1, 7, 6, 3, 8, 2, 4, 9, 5},
		{5, 9, 2, 4, 6, 1, 8, 7, 3},
		{9, 6, 7, 2, 3, 4, 1, 5, 8},
		{2, 1, 8, 5, 7, 9, 3, 6, 4},
		{3, 4, 5, 6, 1, 8, 7, 2, 9},
	}

	actual := grid.GetRow(1)
	expected := board.Row{
		7, 5, 1, 8, 4, 6, 9, 3, 2,
	}

	if actual != expected {
		t.Errorf("GetRow doesn't return expected! Actual: %v", actual)
	}
}

func TestGridGetColumn(t *testing.T) {
	grid := board.Grid{
		{6, 2, 4, 7, 9, 3, 5, 8, 1},
		{7, 5, 1, 8, 4, 6, 9, 3, 2},
		{8, 3, 9, 1, 2, 5, 6, 4, 7},
		{4, 8, 3, 9, 5, 7, 2, 1, 6},
		{1, 7, 6, 3, 8, 2, 4, 9, 5},
		{5, 9, 2, 4, 6, 1, 8, 7, 3},
		{9, 6, 7, 2, 3, 4, 1, 5, 8},
		{2, 1, 8, 5, 7, 9, 3, 6, 4},
		{3, 4, 5, 6, 1, 8, 7, 2, 9},
	}

	actual := grid.GetColumn(1)
	expected := board.Column{
		2, 5, 3, 8, 7, 9, 6, 1, 4,
	}

	if actual != expected {
		t.Errorf("GetRow doesn't return expected!")
	}
}

func TestGridGetBlock(t *testing.T) {
	grid := board.Grid{
		{6, 2, 4, 7, 9, 3, 5, 8, 1},
		{7, 5, 1, 8, 4, 6, 9, 3, 2},
		{8, 3, 9, 1, 2, 5, 6, 4, 7},
		{4, 8, 3, 9, 5, 7, 2, 1, 6},
		{1, 7, 6, 3, 8, 2, 4, 9, 5},
		{5, 9, 2, 4, 6, 1, 8, 7, 3},
		{9, 6, 7, 2, 3, 4, 1, 5, 8},
		{2, 1, 8, 5, 7, 9, 3, 6, 4},
		{3, 4, 5, 6, 1, 8, 7, 2, 9},
	}

	actual := grid.GetBlock(2, 1)
	expected := board.Block{
		{2, 3, 4},
		{5, 7, 9},
		{6, 1, 8},
	}

	if actual != expected {
		t.Errorf("GetBlock doesn't return expected! Actual: %v", actual)
	}
}

func TestGenerateGrid(t *testing.T) {
	grid := board.GenerateGrid()

	for i := 0; i < board.Size; i++ {
		for j := 0; j < board.Size; j++ {
			if grid[i][j] == 0 {
				t.Errorf("grid is not generated completely")
			}
		}
	}
}
