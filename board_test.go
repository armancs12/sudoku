package sudoku_test

import (
	"testing"

	"github.com/serhatscode/sudoku"
)

var (
	board = sudoku.Board{
		{5, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 3, 1, 0, 0, 6, 8, 9, 0},
		{4, 7, 0, 0, 8, 9, 1, 5, 3},
		{0, 0, 8, 7, 0, 0, 0, 0, 0},
		{1, 0, 0, 5, 0, 8, 0, 0, 4},
		{0, 0, 0, 0, 2, 0, 5, 0, 0},
		{9, 6, 5, 8, 3, 0, 0, 2, 1},
		{0, 1, 7, 9, 0, 0, 3, 6, 0},
		{0, 0, 0, 0, 6, 0, 0, 0, 9},
	}
)

func TestColumnHasNumber(t *testing.T) {
	tests := []struct {
		column   int
		number   int
		expected bool
	}{
		{0, 4, true},
		{0, 2, false},
		{1, 7, true},
		{2, 9, false},
		{3, 5, true},
		{8, 4, true},
		{8, 5, false},
	}

	for _, test := range tests {
		actual := board.ColumnHasNumber(test.column, test.number)

		if test.expected != actual {
			t.Errorf("Test: (Column %v has %v should return %v) failed!",
				test.column, test.number, test.expected)
		}
	}
}

func TestRowHasNumber(t *testing.T) {
	tests := []struct {
		row      int
		number   int
		expected bool
	}{
		{0, 5, true},
		{0, 2, false},
		{1, 8, true},
		{2, 2, false},
		{3, 7, true},
		{8, 9, true},
		{8, 4, false},
	}

	for _, test := range tests {
		actual := board.RowHasNumber(test.row, test.number)

		if test.expected != actual {
			t.Errorf("Test: (Row %v has %v should return %v) failed!",
				test.row, test.number, test.expected)
		}
	}
}

func TestBoxHasNumber(t *testing.T) {
	tests := []struct {
		column   int
		row      int
		number   int
		expected bool
	}{
		{0, 0, 7, true},
		{0, 0, 2, false},
		{1, 0, 6, true},
		{1, 0, 2, false},
		{2, 0, 9, true},
		{2, 0, 4, false},
		{0, 1, 8, true},
		{0, 1, 3, false},
		{1, 1, 5, true},
		{0, 2, 4, false},
		{1, 2, 9, true},
		{2, 2, 9, true},
		{2, 2, 4, false},
	}

	for _, test := range tests {
		actual := board.BoxHasNumber(test.column, test.row, test.number)

		if test.expected != actual {
			t.Errorf("Test: (Box (%v, %v) has %v should return %v) failed!",
				test.column, test.row, test.number, test.expected)
		}
	}
}
