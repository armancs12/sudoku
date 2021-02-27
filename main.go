package main

import (
	"fmt"
)

func main() {

	game, err := NewGame()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = game.Start()
	if err != nil {
		fmt.Println(err)
	}
}
