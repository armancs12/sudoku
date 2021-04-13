package game

import "github.com/serhatscode/sudoku/ui"

func DefaultTheme() Theme {
	return Theme{
		Name: "Default",
		Board: ui.BoardTheme{
			Cursor: "#f9da75",
			Border: ui.ColorPair{
				FG: "#344861",
				BG: "#ffffff",
			},
			Cells: ui.BoardCellsTheme{
				Normal: ui.ColorPair{
					FG: "black",
					BG: "#ffffff",
				},
				Predefined: ui.ColorPair{
					FG: "#4a90e2",
					BG: "#ffffff",
				},
				Conflict: ui.ColorPair{
					FG: "#fb3d3f",
					BG: "#f7cfd6",
				},
				Wrong: ui.ColorPair{
					FG: "#fb3d3f",
					BG: "#ffffff",
				},
			},
		},
		Menu: MenuTheme{
			Color: ui.ColorPair{
				FG: "#ffffff",
				BG: "#fb3d3f",
			},
			Cursor: ui.ColorPair{
				FG: "black",
				BG: "#ffffff",
			},
			Box: ui.ColorPair{
				FG: "#ffffff",
				BG: "#fb3d3f",
			},
		},
		Warning: WarningTheme{
			Text: ui.ColorPair{
				FG: "#ffffff",
				BG: "#fb3d3f",
			},
			Box: ui.ColorPair{
				FG: "#ffffff",
				BG: "#fb3d3f",
			},
		},
	}
}

type Theme struct {
	Name    string
	Board   ui.BoardTheme
	Menu    MenuTheme
	Warning WarningTheme
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
