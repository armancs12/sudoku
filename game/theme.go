package game

import (
	"encoding/json"
	"os"
	"path"

	"github.com/serhatsdev/sudoku/ui"
)

type Theme struct {
	Name        string        `json:"name"`
	Board       ui.BoardTheme `json:"board"`
	Menu        ui.ColorPair  `json:"menu"`
	MenuCursor  ui.ColorPair  `json:"menu_cursor"`
	MenuBox     ui.ColorPair  `json:"menu_box"`
	WarningText ui.ColorPair  `json:"warning_text"`
	WarningBox  ui.ColorPair  `json:"warning_box"`
}

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

func CreateThemesFolder() (string, error) {
	themesDir, err := getThemesFolder()
	if err != nil {
		return "", err
	}

	return themesDir, os.MkdirAll(themesDir, os.ModePerm)
}

func LoadThemes() ([]Theme, error) {
	themesDir, err := CreateThemesFolder()
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(themesDir)
	if err != nil {
		return nil, err
	}

	themes := []Theme{DefaultTheme()}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		body, err := os.ReadFile(path.Join(themesDir, file.Name()))
		if err != nil {
			continue
		}

		theme := Theme{}
		json.Unmarshal(body, &theme)

		themes = append(themes, theme)
	}

	return themes, nil
}

func getThemesFolder() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return path.Join(configDir, "sudoku", "themes"), nil
}
