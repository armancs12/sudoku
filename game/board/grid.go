package board

import (
	"math/rand"
	"time"
)

const BlockSize = 3
const Size = BlockSize * BlockSize

type Row [Size]int

type Column [Size]int

type Block [BlockSize][BlockSize]int

type Grid [Size][Size]int

func (grid Grid) GetRow(index int) Row {
	row := Row{}
	for i := 0; i < Size; i++ {
		row[i] = grid[index][i]
	}
	return row
}

func (grid Grid) GetColumn(index int) Column {
	column := Column{}
	for i := 0; i < Size; i++ {
		column[i] = grid[i][index]
	}
	return column
}

func (grid Grid) GetBlock(rowIndex, columnIndex int) Block {
	block := Block{}
	gridRow, gridColumn := blockToGridIndex(rowIndex, columnIndex)

	for i := 0; i < BlockSize; i++ {
		for j := 0; j < BlockSize; j++ {
			block[i][j] = grid[gridRow+i][gridColumn+j]
		}
	}

	return block
}

func blockToGridIndex(blockRowIndex, blockColumnIndex int) (int, int) {
	return blockRowIndex * 3, blockColumnIndex * 3
}

func gridToBlockIndex(gridRowIndex, gridColumnIndex int) (int, int) {
	return int(gridRowIndex / 3), int(gridColumnIndex / 3)
}

func generateFirstRow() Row {
	row := Row{1, 2, 3, 4, 5, 6, 7, 8, 9}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(9, func(i, j int) {
		row[i], row[j] = row[j], row[i]
	})

	return row
}

func findNextEmptyCell(grid *Grid) (int, int, bool) {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if grid[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}

func isRowValidForInsert(row Row, index, value int) bool {
	for i := 0; i < Size; i++ {
		if row[i] == value && index != i {
			return false
		}
	}
	return true
}

func isColumnValidForInsert(column Column, index, value int) bool {
	for i := 0; i < Size; i++ {
		if column[i] == value && index != i {
			return false
		}
	}
	return true
}

func isBlockValidForInsert(block Block, rowIndex, columnIndex, value int) bool {
	for i := 0; i < BlockSize; i++ {
		for j := 0; j < BlockSize; j++ {
			if block[i][j] == value && rowIndex != i && columnIndex != j {
				return false
			}
		}
	}
	return true
}

func isGridValidForInsert(grid *Grid, rowIndex, columnIndex, value int) bool {
	blockRow := rowIndex % Size
	blockColumn := columnIndex % Size

	return isRowValidForInsert(grid.GetRow(rowIndex), columnIndex, value) &&
		isColumnValidForInsert(grid.GetColumn(columnIndex), rowIndex, value) &&
		isBlockValidForInsert(
			grid.GetBlock(gridToBlockIndex(rowIndex, columnIndex)),
			blockRow,
			blockColumn,
			value,
		)
}

func completeGrid(grid *Grid) *Grid {
	rowIndex, columnIndex, exist := findNextEmptyCell(grid)
	if !exist {
		return grid
	}

	for i := 1; i <= Size; i++ {
		if isGridValidForInsert(grid, rowIndex, columnIndex, i) {
			grid[rowIndex][columnIndex] = i

			if solved := completeGrid(grid); solved != nil {
				return solved
			}
		}
		grid[rowIndex][columnIndex] = 0
	}

	return nil
}

func GenerateGrid() Grid {
	grid := Grid{generateFirstRow()}
	return *completeGrid(&grid)
}
