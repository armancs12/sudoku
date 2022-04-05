package main

import (
	"fmt"

	"github.com/serhatsdev/sudoku/game"
	"github.com/serhatsdev/sudoku/game/ui"
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
