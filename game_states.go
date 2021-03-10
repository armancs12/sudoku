package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/serhatscode/sudoku/board"
	"github.com/serhatscode/sudoku/ui"
)

// State is an interface for game states
// The current game state will handle:
//
// - What to do when game resized
// - What to do when key pressed
// - What to draw
type State interface {
	// OnResize will be invoked when terminal resized
	OnResize(event *tcell.EventResize)
	// OnKeyPress will be invoked when key pressed
	OnKeyPress(event *tcell.EventKey)
	// Draw will be invoked after every event
	Draw()
}

type playState struct {
	Game Game
	Pos  board.Point2
}

// NewPlayState returns a new play state
func NewPlayState(game Game) State {
	return &playState{
		Game: game,
		// center of the board
		Pos: board.Point2{X: 4, Y: 4},
	}
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

type menuState struct {
	Game      Game
	Pos       int
	Options   []string
	Functions []func()
}

// NewMenuState returns a new menu state
func NewMenuState(game Game) State {
	return &menuState{
		Game: game,
		Pos:  0,
		Options: []string{
			"Resume",
			"New Game",
			"Exit",
		},
		Functions: []func(){
			func() {
				game.PopState()
			},
			func() {
				game.SetBoard(board.New(board.Medium))
				game.PopState()
			},
			func() {
				game.Exit()
			},
		},
	}
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

type smallSizeState struct {
	Game   Game
	Width  int
	Height int
}

// NewSmallSizeState returns a new small size state
func NewSmallSizeState(game Game, width, height int) State {
	return &smallSizeState{game, width, height}
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
