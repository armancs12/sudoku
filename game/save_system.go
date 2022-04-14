package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/serhatsdev/sudoku/game/board"
	"github.com/serhatsdev/sudoku/game/theme"
)

var ErrNoSavedGame = errors.New("there is no saved game")

var ErrSaveCorrupted = errors.New("save file is corrupted")

type SaveData struct {
	Board board.Board
	Theme theme.Theme
}

type SaveDataJSON struct {
	Version   string `json:"version"`
	BoardData string `json:"board_data"`
	ThemeName string `json:"theme_name"`
}

func boolToInt(value bool) int {
	if value {
		return 1
	}

	return 0
}

func intToBool(value int) bool {
	if value > 0 {
		return true
	}

	return false
}

func getBoardData(b board.Board) string {
	valueData := ""
	correctData := ""
	predefinedData := ""

	for i := 0; i < board.Size; i++ {
		for j := 0; j < board.Size; j++ {
			pos := board.Point2{X: j, Y: i}

			valueData += fmt.Sprint(b.Get(pos))
			correctData += fmt.Sprint(b.GetCorrect(pos))
			predefinedData += fmt.Sprint(boolToInt(b.IsPredefined(pos)))
		}
	}

	return strings.Join([]string{valueData, correctData, predefinedData}, "-")
}

func getGridFromStringData(data string) (board.Grid, error) {
	grid := board.Grid{}
	for i := 0; i < len(data); i++ {
		rowIndex := int(i / board.Size)
		columnIndex := i % board.Size

		value, err := strconv.Atoi(string(data[i]))
		if err != nil {
			return board.Grid{}, ErrSaveCorrupted
		}

		grid[rowIndex][columnIndex] = value
	}

	return grid, nil
}

func getPredefinedGridFromStringData(data string) ([board.Size][board.Size]bool, error) {
	grid := [board.Size][board.Size]bool{}
	for i := 0; i < len(data); i++ {
		rowIndex := int(i / board.Size)
		columnIndex := i % board.Size

		value, err := strconv.Atoi(string(data[i]))
		if err != nil {
			return [board.Size][board.Size]bool{}, ErrSaveCorrupted
		}

		grid[rowIndex][columnIndex] = intToBool(value)
	}

	return grid, nil
}

func loadBoard(boardData string) (board.Board, error) {
	datas := strings.Split(boardData, "-")
	if len(datas) != 3 {
		return nil, ErrSaveCorrupted
	}

	uncompleteGrid, err := getGridFromStringData(datas[0])
	if err != nil {
		return nil, err
	}
	completeGrid, err := getGridFromStringData(datas[1])
	if err != nil {
		return nil, err
	}
	predefinedGrid, err := getPredefinedGridFromStringData(datas[2])
	if err != nil {
		return nil, err
	}

	return board.NewCustom(uncompleteGrid, completeGrid, predefinedGrid), nil
}

func getFirstThemeByNameOrDefault(themes []theme.Theme, name string) theme.Theme {
	for _, theme := range themes {
		if theme.Name == name {
			return theme
		}
	}

	return themes[0]
}

func getSaveFile() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return path.Join(configDir, "sudoku", "save.json"), nil
}

func LoadSavedGame() (SaveData, error) {
	saveFile, err := getSaveFile()
	if err != nil {
		return SaveData{}, err
	}

	savedatajson := SaveDataJSON{}
	file, err := os.ReadFile(saveFile)
	json.Unmarshal(file, &savedatajson)

	themes, err := theme.GetThemes()
	if err != nil {
		return SaveData{}, err
	}

	board, err := loadBoard(savedatajson.BoardData)
	if err != nil {
		return SaveData{}, err
	}
	theme := getFirstThemeByNameOrDefault(themes, savedatajson.ThemeName)

	savedata := SaveData{
		Board: board,
		Theme: theme,
	}

	return savedata, nil
}

func SaveGame(savedata SaveData) error {
	saveFile, err := getSaveFile()
	if err != nil {
		return err
	}

	savedatajson := SaveDataJSON{
		Version:   Version,
		ThemeName: savedata.Theme.Name,
		BoardData: getBoardData(savedata.Board),
	}

	data, err := json.Marshal(savedatajson)
	if err != nil {
		return err
	}

	return os.WriteFile(saveFile, data, os.ModePerm)
}
