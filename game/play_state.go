package game

import (
	"strconv"

	"github.com/serhatsdev/sudoku/game/board"
	"github.com/serhatsdev/sudoku/game/ui"
)

type playState struct {
	Game Game
	Pos  board.Point2
}

func (ps *playState) OnResize(width, height int) {
	if width < ps.Game.MinWidth() || height < ps.Game.MinHeight() {
		ps.Game.PushState(NewSmallSizeState(ps.Game, width, height))
	}
}

func (ps *playState) OnKeyPress(key string) {
	if key == "esc" {
		ps.Game.PushState(NewMenuState(ps.Game))
		return
	}

	if key == "arrow_up" && ps.Pos.Y > 0 {
		ps.Pos.Y--
	} else if key == "arrow_down" && ps.Pos.Y < 8 {
		ps.Pos.Y++
	} else if key == "arrow_left" && ps.Pos.X > 0 {
		ps.Pos.X--
	} else if key == "arrow_right" && ps.Pos.X < 8 {
		ps.Pos.X++
	} else {
		if key == "e" || key == "E" {
			ps.Game.Board().Set(ps.Pos, 0)
		} else {
			num, err := strconv.Atoi(key)
			if err == nil && num > 0 && num < 10 {
				ps.Game.Board().Set(ps.Pos, board.Value(num))
			}
		}
	}
}

func (ps *playState) Draw() {
	ps.Game.Client().DrawCenter(&ui.BoardWidget{
		Board:     ps.Game.Board(),
		CursorPos: ps.Pos,
		Theme:     ps.Game.Theme().Board,
	})
}
