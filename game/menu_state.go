package game

import (
	"github.com/gdamore/tcell/v2"
	"github.com/serhatscode/sudoku/ui"
)

type menuState struct {
	Game      Game
	Pos       int
	Options   []string
	Functions []func()
}

func (ms *menuState) OnResize(event *tcell.EventResize) {
	width, height := event.Size()
	if width < ms.Game.MinWidth() || height < ms.Game.MinHeight() {
		ms.Game.PushState(NewSmallSizeState(ms.Game, width, height))
	}
}

func (ms *menuState) OnKeyPress(event *tcell.EventKey) {
	key := event.Key()
	if key == tcell.KeyESC {
		ms.Game.PopState()
		return
	}

	if key == tcell.KeyUp {
		ms.Pos = (len(ms.Options) + ms.Pos - 1) % len(ms.Options)
	} else if key == tcell.KeyDown {
		ms.Pos = (ms.Pos + 1) % len(ms.Options)
	} else if key == tcell.KeyEnter {
		ms.Functions[ms.Pos]()
	}
}

func (ms *menuState) Draw() {
	ui.DrawCenter(&ui.BoxWidget{
		Child: &ui.MenuWidget{
			Options:          ms.Options,
			CursorIndex:      ms.Pos,
			AlignCenter:      true,
			Color:            tcell.ColorWhite,
			Background:       tcell.ColorRed,
			CursorColor:      tcell.ColorWhite,
			CursorBackground: tcell.ColorBlack,
		},
		PaddingTop:    1,
		PaddingBottom: 1,
		PaddingLeft:   1,
		PaddingRight:  1,
		Color:         tcell.ColorWhite,
		Background:    tcell.ColorRed,
	})
}
