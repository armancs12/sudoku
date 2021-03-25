package game

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/serhatscode/sudoku/ui"
)

type smallSizeState struct {
	Game   Game
	Width  int
	Height int
}

func (sss *smallSizeState) OnResize(event *tcell.EventResize) {
	width, height := event.Size()
	if width >= sss.Game.MinWidth() && height >= sss.Game.MinHeight() {
		sss.Game.PopState()
	} else {
		sss.Width, sss.Height = width, height
	}
}

func (sss *smallSizeState) OnKeyPress(event *tcell.EventKey) {}

func (sss *smallSizeState) Draw() {
	message := fmt.Sprintf("Please resize to\n at least %dx%d",
		sss.Game.MinWidth(), sss.Game.MinHeight())
	current := fmt.Sprintf("%dx%d", sss.Width, sss.Height)

	ui.DrawUpRightCorner(&ui.TextWidget{String: current})
	ui.DrawCenter(&ui.BoxWidget{
		Child: &ui.TextWidget{
			String:      message,
			AlignCenter: true,
			Color:       tcell.ColorWhite,
			Background:  tcell.ColorRed,
		},
		Fill:          true,
		PaddingTop:    1,
		PaddingBottom: 1,
		PaddingLeft:   1,
		PaddingRight:  1,
		Color:         tcell.ColorWhite,
		Background:    tcell.ColorRed,
	})
}
