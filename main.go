package main

import (
	"fmt"

	"github.com/serhatscode/sudoku/game"
)

func main() {

	game, err := game.NewGame()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = game.Start()
	if err != nil {
		fmt.Println(err)
	}
}
