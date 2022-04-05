package game

import (
	"github.com/serhatsdev/sudoku/game/ui"
)

type menuOption struct {
	title    string
	function func()
}

type menuState struct {
	Game    Game
	Pos     int
	Options []menuOption
}

func (ms *menuState) OnResize(width, height int) {
	if width < ms.Game.MinWidth() || height < ms.Game.MinHeight() {
		ms.Game.PushState(NewSmallSizeState(ms.Game, width, height))
	}
}

func (ms *menuState) OnKeyPress(key string) {
	if key == "esc" {
		ms.Game.PopState()
		return
	}

	if key == "arrow_up" {
		ms.Pos = (len(ms.Options) + ms.Pos - 1) % len(ms.Options)
	} else if key == "arrow_down" {
		ms.Pos = (ms.Pos + 1) % len(ms.Options)
	} else if key == "enter" {
		ms.Options[ms.Pos].function()
	}
}

func (ms *menuState) Draw() {
	ms.Game.Client().DrawCenter(&ui.BoxWidget{
		Child: &ui.MenuWidget{
			Options:     getTitlesFromOptions(ms.Options),
			CursorIndex: ms.Pos,
			HAlign:      ui.HAlignCenter,
			Color:       ms.Game.Theme().Menu,
			Cursor:      ms.Game.Theme().MenuCursor,
		},
		PaddingTop:    1,
		PaddingBottom: 1,
		PaddingLeft:   1,
		PaddingRight:  1,
		Color:         ms.Game.Theme().MenuBox,
	})
}

func getTitlesFromOptions(options []menuOption) []string {
	titles := []string{}
	for i := 0; i < len(options); i++ {
		titles = append(titles, options[i].title)
	}
	return titles
}
