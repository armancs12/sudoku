package game

import (
	"github.com/gdamore/tcell/v2"
	"github.com/serhatscode/sudoku/board"
	"github.com/serhatscode/sudoku/ui"
)

type playState struct {
	Game Game
	Pos  board.Point2
}

func (ps *playState) OnResize(event *tcell.EventResize) {
	width, height := event.Size()
	if width < ps.Game.MinWidth() || height < ps.Game.MinHeight() {
		ps.Game.PushState(NewSmallSizeState(ps.Game, width, height))
	}
}

func (ps *playState) OnKeyPress(event *tcell.EventKey) {
	key := event.Key()

	if key == tcell.KeyESC {
		ps.Game.PushState(NewMenuState(ps.Game))
		return
	}

	if key == tcell.KeyUp && ps.Pos.Y > 0 {
		ps.Pos.Y--
	} else if key == tcell.KeyDown && ps.Pos.Y < 8 {
		ps.Pos.Y++
	} else if key == tcell.KeyLeft && ps.Pos.X > 0 {
		ps.Pos.X--
	} else if key == tcell.KeyRight && ps.Pos.X < 8 {
		ps.Pos.X++
	} else {
		char := event.Rune()
		if char == 'e' || char == 'E' {
			ps.Game.Board().Set(ps.Pos, 0)
		} else {
			num := int(char - '0')
			if num > 0 && num < 10 {
				ps.Game.Board().Set(ps.Pos, board.Value(num))
			}
		}
	}
}

func (ps *playState) Draw() {
	ui.DrawCenter(&ui.BoardWidget{
		Board:        ps.Game.Board(),
		CursorPos:    ps.Pos,
		CursorBG:     tcell.ColorRed,
		NumberFG:     tcell.ColorBlack,
		BorderFG:     tcell.ColorBlack,
		PredefinedFG: tcell.ColorBlue,
		ConflictFG:   tcell.ColorRed,
		Background:   tcell.ColorWhite,
	})
}
