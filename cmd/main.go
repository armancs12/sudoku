package main

import (
	"fmt"

	"github.com/serhatscode/sudoku"
)

func main() {

	game, err := sudoku.NewGame()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = game.Start()
	if err != nil {
		fmt.Println(err)
	}
}
