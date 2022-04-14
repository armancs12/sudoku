package theme

import (
	"encoding/json"
	"os"
	"path"
	"strings"
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

var loadedThemes []Theme

func getDefaultThemes() map[string]Theme {
	defaultLight := Theme{
		Name: "Default Light",
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

	return map[string]Theme{
		"default_light.json": defaultLight,
	}
}

func getThemesDirectory() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return path.Join(configDir, "sudoku", "themes"), nil
}

func isThemesDirExist() (bool, error) {
	themesDir, err := getThemesDirectory()
	if err != nil {
		return false, err
	}

	_, err = os.Stat(themesDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		}

		return false, err
	}

	return true, nil
}

func getFileType(filename string) string {
	parts := strings.Split(filename, ".")
	return parts[len(parts)-1:][0]
}

func getThemeJsonFiles() ([]string, error) {
	themesDir, err := getThemesDirectory()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(themesDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}

		return nil, err
	}

	jsonFiles := []string{}
	for _, entry := range entries {
		if !entry.IsDir() && getFileType(entry.Name()) == "json" {
			jsonFiles = append(jsonFiles, path.Join(themesDir, entry.Name()))
		}
	}

	return jsonFiles, nil
}

func loadThemes() ([]Theme, error) {
	jsonFiles, err := getThemeJsonFiles()
	if err != nil {
		return nil, err
	}
	themes := []Theme{}

	for _, file := range jsonFiles {
		data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		theme := Theme{}
		err = json.Unmarshal([]byte(data), &theme)
		if err != nil {
			return nil, err
		}

		themes = append(themes, theme)
	}

	return themes, nil
}

func createDefaultThemes() ([]Theme, error) {
	themesDir, err := getThemesDirectory()
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(themesDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	themesWithFilename := getDefaultThemes()
	themes := []Theme{}

	for filename, theme := range themesWithFilename {
		data, err := json.Marshal(theme)
		if err != nil {
			return nil, err
		}

		err = os.WriteFile(path.Join(themesDir, filename), data, os.ModePerm)
		if err != nil {
			return nil, err
		}

		themes = append(themes, theme)
	}

	return themes, nil
}

func GetThemes() ([]Theme, error) {
	if loadedThemes != nil && len(loadedThemes) > 0 {
		return loadedThemes, nil
	}

	themes, err := loadThemes()
	if err != nil {
		return nil, err
	}

	if len(themes) == 0 {
		themes, err = createDefaultThemes()
		if err != nil {
			return nil, err
		}
	}

	loadedThemes = themes
	return themes, nil
}
