package theme

import (
	"encoding/json"
	"os"
	"path"
)

type Theme struct {
	Name        string     `json:"name"`
	Board       BoardTheme `json:"board"`
	Menu        ColorPair  `json:"menu"`
	MenuCursor  ColorPair  `json:"menu_cursor"`
	MenuBox     ColorPair  `json:"menu_box"`
	WarningText ColorPair  `json:"warning_text"`
	WarningBox  ColorPair  `json:"warning_box"`
}

type ColorPair struct {
	FG string `json:"fg"`
	BG string `json:"bg"`
}

type BoardCellsTheme struct {
	Normal     ColorPair `json:"normal"`
	Predefined ColorPair `json:"predefined"`
	Conflict   ColorPair `json:"conflict"`
	Wrong      ColorPair `json:"wrong"`
}

type BoardTheme struct {
	Cursor string          `json:"cursor"`
	Border ColorPair       `json:"border"`
	Cells  BoardCellsTheme `json:"cells"`
}

func Default() Theme {
	return Theme{
		Name: "Default",
		Board: BoardTheme{
			Cursor: "#f9da75",
			Border: ColorPair{
				FG: "#344861",
				BG: "#ffffff",
			},
			Cells: BoardCellsTheme{
				Normal: ColorPair{
					FG: "black",
					BG: "#ffffff",
				},
				Predefined: ColorPair{
					FG: "#4a90e2",
					BG: "#ffffff",
				},
				Conflict: ColorPair{
					FG: "#fb3d3f",
					BG: "#f7cfd6",
				},
				Wrong: ColorPair{
					FG: "#fb3d3f",
					BG: "#ffffff",
				},
			},
		},
		Menu: ColorPair{
			FG: "#ffffff",
			BG: "#fb3d3f",
		},
		MenuCursor: ColorPair{
			FG: "black",
			BG: "#ffffff",
		},
		MenuBox: ColorPair{
			FG: "#ffffff",
			BG: "#fb3d3f",
		},
		WarningText: ColorPair{
			FG: "#ffffff",
			BG: "#fb3d3f",
		},
		WarningBox: ColorPair{
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

func Load() ([]Theme, error) {
	themesDir, err := CreateThemesFolder()
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(themesDir)
	if err != nil {
		return nil, err
	}

	themes := []Theme{Default()}
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
