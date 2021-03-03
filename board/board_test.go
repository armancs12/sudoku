package board_test

import (
	"testing"

	"github.com/serhatscode/sudoku/board"
)

func getBoard() board.IBoard {
	return board.NewCustom(
		[9][9]board.Value{
			{0, 2, 0, 0, 9, 0, 5, 8, 0},
			{7, 5, 0, 8, 4, 0, 9, 3, 2},
			{8, 0, 9, 1, 2, 0, 0, 4, 0},
			{4, 0, 0, 0, 5, 0, 2, 1, 6},
			{0, 7, 6, 3, 0, 2, 0, 0, 5},
			{5, 0, 2, 0, 0, 0, 8, 7, 0},
			{0, 6, 0, 0, 3, 4, 1, 0, 8},
			{2, 1, 8, 5, 0, 9, 0, 0, 4},
			{3, 4, 0, 0, 0, 8, 7, 2, 0},
		},
		[9][9]board.Value{
			{6, 2, 4, 7, 9, 3, 5, 8, 1},
			{7, 5, 1, 8, 4, 6, 9, 3, 2},
			{8, 3, 9, 1, 2, 5, 6, 4, 7},
			{4, 8, 3, 9, 5, 7, 2, 1, 6},
			{1, 7, 6, 3, 8, 2, 4, 9, 5},
			{5, 9, 2, 4, 6, 1, 8, 7, 3},
			{9, 6, 7, 2, 3, 4, 1, 5, 8},
			{2, 1, 8, 5, 7, 9, 3, 6, 4},
			{3, 4, 5, 6, 1, 8, 7, 2, 9},
		},
	)
}

func TestGet(t *testing.T) {
	tBoard := getBoard()

	tests := []struct {
		pos      board.Point2
		expected board.Value
	}{
		{
			pos:      board.Point2{0, 0},
			expected: 0,
		},
		{
			pos:      board.Point2{8, 8},
			expected: 0,
		},
		{
			pos:      board.Point2{0, 8},
			expected: 3,
		},
		{
			pos:      board.Point2{8, 0},
			expected: 0,
		},
		{
			pos:      board.Point2{1, 1},
			expected: 5,
		},
	}

	for _, test := range tests {
		actual := tBoard.Get(test.pos)
		if test.expected != actual {
			t.Errorf("board.Get(%d,%d) failed: Expected: %d, Actual:%d",
				test.pos.X, test.pos.Y, test.expected, actual)
		}
	}
}

func TestSet(t *testing.T) {
	tBoard := getBoard()

	tests := []struct {
		pos      board.Point2
		value    board.Value
		expected board.Value
	}{
		{
			pos:      board.Point2{0, 0},
			value:    1,
			expected: 1,
		},
		{
			pos:      board.Point2{8, 8},
			value:    2,
			expected: 2,
		},
		{
			pos:      board.Point2{0, 8},
			value:    9,
			expected: 3,
		},
		{
			pos:      board.Point2{8, 0},
			value:    9,
			expected: 9,
		},
		{
			pos:      board.Point2{1, 1},
			value:    2,
			expected: 5,
		},
	}

	for _, test := range tests {
		tBoard.Set(test.pos, test.value)
		actual := tBoard.Get(test.pos)
		if test.expected != actual {
			t.Errorf("board.Set(%d,%d) failed: Expected: %d, Actual:%d",
				test.pos.X, test.pos.Y, test.expected, actual)
		}
	}
}

func TestIsPredefined(t *testing.T) {
	tBoard := getBoard()
	tBoard.Set(board.Point2{0, 0}, 2)
	tBoard.Set(board.Point2{8, 0}, 1)
	tBoard.Set(board.Point2{8, 8}, 9)

	tests := []struct {
		pos      board.Point2
		expected bool
	}{
		{
			pos:      board.Point2{0, 0},
			expected: false,
		},
		{
			pos:      board.Point2{8, 8},
			expected: false,
		},
		{
			pos:      board.Point2{0, 8},
			expected: true,
		},
		{
			pos:      board.Point2{8, 0},
			expected: false,
		},
		{
			pos:      board.Point2{1, 1},
			expected: true,
		},
	}

	for _, test := range tests {
		actual := tBoard.IsPredefined(test.pos)
		if test.expected != actual {
			t.Errorf("board.IsPredefined(%d,%d) failed: Expected: %v, Actual:%v",
				test.pos.X, test.pos.Y, test.expected, actual)
		}
	}
}

func TestIsCorrect(t *testing.T) {
	tBoard := getBoard()
	tBoard.Set(board.Point2{0, 0}, 2)
	tBoard.Set(board.Point2{8, 0}, 1)
	tBoard.Set(board.Point2{8, 8}, 9)

	tests := []struct {
		pos      board.Point2
		expected bool
	}{
		{
			pos:      board.Point2{0, 0},
			expected: false,
		},
		{
			pos:      board.Point2{8, 8},
			expected: true,
		},
		{
			pos:      board.Point2{0, 8},
			expected: true,
		},
		{
			pos:      board.Point2{8, 0},
			expected: true,
		},
		{
			pos:      board.Point2{1, 1},
			expected: true,
		},
	}

	for _, test := range tests {
		actual := tBoard.IsCorrect(test.pos)
		if test.expected != actual {
			t.Errorf("board.IsCorrect(%d,%d) failed: Expected: %v, Actual:%v",
				test.pos.X, test.pos.Y, test.expected, actual)
		}
	}
}

func TestGetConflicts(t *testing.T) {
	tBoard := getBoard()

	tests := []struct {
		pos      board.Point2
		value    board.Value
		expected []board.Point2
	}{
		{
			pos:   board.Point2{0, 0},
			value: 2,
			expected: []board.Point2{
				{1, 0}, {0, 7},
			},
		},
		{
			pos:      board.Point2{8, 8},
			value:    9,
			expected: []board.Point2{},
		},
		{
			pos:   board.Point2{8, 0},
			value: 9,
			expected: []board.Point2{
				{4, 0}, {6, 1},
			},
		},
		{
			pos:   board.Point2{0, 8},
			value: 5,
			expected: []board.Point2{
				{0, 5},
			},
		},
		{
			pos:   board.Point2{1, 3},
			value: 6,
			expected: []board.Point2{
				{8, 3}, {2, 4}, {1, 6},
			},
		},
	}

	for _, test := range tests {
		actual := tBoard.GetConflicts(test.pos, test.value)
		if !equals(actual, test.expected) {
			t.Errorf("board.GetConflicts(%d,%d, %d) failed! Map: %v",
				test.pos.X, test.pos.Y, test.value, actual)
		}
	}
}

func TestGetPositions(t *testing.T) {
	tBoard := getBoard()

	tests := []struct {
		value    board.Value
		expected []board.Point2
	}{
		{
			value: 2,
			expected: []board.Point2{
				{1, 0},
				{8, 1},
				{4, 2},
				{6, 3},
				{5, 4},
				{2, 5},
				{0, 7},
				{7, 8},
			},
		},
		{
			value: 9,
			expected: []board.Point2{
				{4, 0}, {6, 1}, {2, 2}, {5, 7},
			},
		},
		{
			value: 0,
			expected: []board.Point2{
				{0, 0}, {2, 0}, {3, 0}, {5, 0}, {8, 0},
				{2, 1}, {5, 1},
				{1, 2}, {5, 2}, {6, 2}, {8, 2},
				{1, 3}, {2, 3}, {3, 3}, {5, 3},
				{0, 4}, {4, 4}, {6, 4}, {7, 4},
				{1, 5}, {3, 5}, {4, 5}, {5, 5}, {8, 5},
				{0, 6}, {2, 6}, {3, 6}, {7, 6},
				{4, 7}, {6, 7}, {7, 7},
				{2, 8}, {3, 8}, {4, 8}, {8, 8},
			},
		},
	}

	for _, test := range tests {
		actual := tBoard.GetPositions(test.value)
		if !equals(actual, test.expected) {
			t.Errorf("board.GetPositions(%d) failed! Map: %v", test.value, actual)
		}
	}
}

// ================== util functions =================
func equals(pMap map[board.Point2]struct{}, pSlice []board.Point2) bool {
	if len(pMap) != len(pSlice) {
		return false
	}

	for _, point := range pSlice {
		_, contains := pMap[point]
		if !contains {
			return false
		}
	}

	return true
}