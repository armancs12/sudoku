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
		Menu: ui.ColorPair{
			FG: "#ffffff",
			BG: "#fb3d3f",
		},
		MenuCursor: ui.ColorPair{
			FG: "black",
			BG: "#ffffff",
		},
		MenuBox: ui.ColorPair{
			FG: "#ffffff",
			BG: "#fb3d3f",
		},
		WarningText: ui.ColorPair{
			FG: "#ffffff",
			BG: "#fb3d3f",
		},
		WarningBox: ui.ColorPair{
			FG: "#ffffff",
			BG: "#fb3d3f",
		},
	}
}

type Theme struct {
	Name        string
	Board       ui.BoardTheme
	Menu        ui.ColorPair
	MenuCursor  ui.ColorPair
	MenuBox     ui.ColorPair
	WarningText ui.ColorPair
	WarningBox  ui.ColorPair
}
