package game

import "github.com/serhatscode/sudoku/ui"

type Theme struct {
	Name    string
	Board   BoardTheme
	Menu    MenuTheme
	Warning WarningTheme
}

type BoardTheme struct {
	Cursor string
	Border ui.ColorPair
	Cells  BoardCellColors
}

type BoardCellColors struct {
	Normal     ui.ColorPair
	Predefined ui.ColorPair
	Conflict   ui.ColorPair
	Wrong      ui.ColorPair
}

type MenuTheme struct {
	Color  ui.ColorPair
	Cursor ui.ColorPair
	Box    ui.ColorPair
}

type WarningTheme struct {
	Box  ui.ColorPair
	Text ui.ColorPair
}
