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

func TestIsGuessValid(t *testing.T) {
	tests := []struct {
		column   int
		row      int
		guess    int
		expected bool
	}{
		{1, 0, 2, true},
		{1, 0, 4, false},
		{0, 1, 3, false},
		{0, 1, 9, false},
		{3, 2, 1, false},
		{3, 2, 2, true},
		{3, 2, 3, false},
		{3, 2, 4, false},
		{3, 2, 5, false},
		{3, 2, 6, false},
		{3, 2, 7, false},
		{3, 2, 8, false},
		{3, 2, 9, false},
		{7, 8, 2, false},
		{7, 8, 4, true},
	}

	for _, test := range tests {
		actual := board.IsGuessValid(test.column, test.row, test.guess)

		if test.expected != actual {
			t.Errorf("Test: (Guess %v in (%v, %v) is valid should return %v) failed!",
				test.guess, test.column, test.row, test.expected)
		}
	}
}

func TestGetNextEmptyCell(t *testing.T) {
	tests := []struct {
		board         sudoku.Board
		expectedCol   int
		expectedRow   int
		expectedFound bool
	}{
		{
			board: sudoku.Board{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{6, 5, 4, 9, 8, 7, 1, 2, 3},
				{9, 8, 7, 1, 2, 3, 4, 5, 6},
				{3, 1, 2, 6, 4, 5, 9, 7, 8},
				{5, 4, 6, 8, 7, 9, 3, 1, 2},
				{7, 9, 8, 3, 1, 2, 6, 4, 5},
				{2, 3, 1, 5, 0, 4, 8, 9, 7},
				{4, 6, 5, 7, 9, 8, 2, 3, 1},
				{8, 7, 9, 2, 3, 1, 5, 6, 4},
			},
			expectedCol:   4,
			expectedRow:   6,
			expectedFound: true,
		},
		{
			board: sudoku.Board{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{6, 5, 4, 9, 8, 7, 1, 2, 3},
				{9, 8, 7, 1, 2, 3, 4, 5, 6},
				{3, 1, 2, 6, 4, 0, 9, 7, 8},
				{5, 4, 6, 8, 7, 9, 3, 1, 2},
				{7, 9, 8, 0, 1, 2, 6, 4, 5},
				{2, 3, 1, 5, 6, 4, 8, 9, 7},
				{4, 6, 5, 7, 9, 8, 2, 3, 1},
				{8, 7, 9, 2, 3, 1, 5, 6, 4},
			},
			expectedCol:   5,
			expectedRow:   3,
			expectedFound: true,
		},
		{
			board: sudoku.Board{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{6, 5, 4, 9, 8, 7, 1, 2, 3},
				{9, 8, 7, 1, 2, 3, 4, 5, 6},
				{3, 1, 2, 6, 4, 5, 0, 7, 8},
				{5, 4, 6, 8, 7, 9, 3, 1, 2},
				{0, 9, 8, 3, 1, 2, 6, 4, 5},
				{2, 3, 1, 5, 6, 4, 8, 9, 7},
				{4, 6, 5, 7, 9, 8, 2, 3, 1},
				{8, 7, 9, 2, 3, 1, 5, 6, 4},
			},
			expectedCol:   6,
			expectedRow:   3,
			expectedFound: true,
		},
		{
			board: sudoku.Board{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{6, 5, 4, 9, 8, 7, 1, 2, 3},
				{9, 8, 7, 1, 2, 3, 4, 5, 6},
				{3, 1, 2, 6, 4, 5, 9, 7, 8},
				{5, 4, 6, 8, 7, 9, 3, 1, 2},
				{7, 9, 8, 3, 1, 2, 6, 4, 5},
				{2, 3, 1, 5, 6, 4, 8, 9, 7},
				{4, 6, 5, 7, 9, 8, 2, 3, 1},
				{8, 7, 9, 2, 3, 1, 5, 6, 4},
			},
			expectedCol:   0,
			expectedRow:   0,
			expectedFound: false,
		},
	}

	for _, test := range tests {
		actualCol, actualRow, actualFound := test.board.GetNextEmptyCell()

		if test.expectedFound != actualFound ||
			test.expectedCol != actualCol ||
			test.expectedRow != actualRow {
			t.Errorf("Expected (%v, %v, %v) but got (%v, %v, %v)!",
				test.expectedCol, test.expectedRow, test.expectedFound,
				actualCol, actualRow, actualFound)
		}
	}
}
