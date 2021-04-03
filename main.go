package main

import (
	"fmt"

	"github.com/serhatscode/sudoku/game"
	"github.com/serhatscode/sudoku/ui"
)

func main() {
	client, err := ui.NewTCellClient()
	checkErr(err)

	game, err := game.NewGame(client)
	checkErr(err)

	err = game.Start()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("error:", err.Error())
	}
}
