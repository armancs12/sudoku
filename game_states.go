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
// - What to do when k	ey pressed
// - What to draw
type State interface {
	// OnResize will be invoked when terminal resized
	OnResize(event *tcell.EventResize)
	// OnKeyPress will be invoked when key pressed
	OnKeyPress(event *tcell.EventKey)
	// Draw will be invoked after every event
	Draw()
}

// PlayState is the state of gameplay
type PlayState struct {
	Game       IGame
	PosX, PosY int
}

// NewPlayState returns a new PlayState
func NewPlayState(game IGame) *PlayState {
	return &PlayState{
		Game: game,
		// center of the board
		PosX: 4,
		PosY: 4,
	}
}

// OnResize checks if size is big enough
func (ps *PlayState) OnResize(event *tcell.EventResize) {
	width, height := event.Size()
	if width < ps.Game.MinWidth() || height < ps.Game.MinHeight() {
		ps.Game.PushState(NewSmallSizeState(ps.Game, width, height))
	}
}

// OnKeyPress changes the board indicator position
func (ps *PlayState) OnKeyPress(event *tcell.EventKey) {
	key := event.Key()

	if key == tcell.KeyESC {
		ps.Game.PushState(NewMenuState(ps.Game))
		return
	}

	if key == tcell.KeyUp && ps.PosY > 0 {
		ps.PosY--
	} else if key == tcell.KeyDown && ps.PosY < 8 {
		ps.PosY++
	} else if key == tcell.KeyLeft && ps.PosX > 0 {
		ps.PosX--
	} else if key == tcell.KeyRight && ps.PosX < 8 {
		ps.PosX++
	} else {
		char := event.Rune()
		if char == 'e' || char == 'E' {
			ps.Game.Board()[ps.PosY][ps.PosX] = 0
		} else {
			num := int(char - '0')
			if num > 0 && num < 10 {
				ps.Game.Board()[ps.PosY][ps.PosX] = num
			}
		}
	}
}

// Draw draws the board
func (ps *PlayState) Draw() {
	ui.DrawCenter(&ui.BoardWidget{
		Board:       *ps.Game.Board(),
		CursorX:     ps.PosX,
		CursorY:     ps.PosY,
		CursorColor: tcell.ColorRed,
		NumberColor: tcell.ColorBlack,
		BorderColor: tcell.ColorBlack,
		Background:  tcell.ColorWhite,
	})
}

// MenuState is the state of game menu
type MenuState struct {
	Game      IGame
	Pos       int
	Options   []string
	Functions []func()
}

// NewMenuState returns a new MenuState
func NewMenuState(game IGame) *MenuState {
	return &MenuState{
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
				newBoard := board.NewBoard(40)
				game.SetBoard(&newBoard)
				game.PopState()
			},
			func() {
				game.Exit()
			},
		},
	}
}

// OnResize checks if size is big enough
func (ms *MenuState) OnResize(event *tcell.EventResize) {
	width, height := event.Size()
	if width < ms.Game.MinWidth() || height < ms.Game.MinHeight() {
		ms.Game.PushState(NewSmallSizeState(ms.Game, width, height))
	}
}

// OnKeyPress changes the menu indicator position
func (ms *MenuState) OnKeyPress(event *tcell.EventKey) {
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

// Draw draws the game menu
func (ms *MenuState) Draw() {
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

// SmallSizeState is the state of terminal size not being big enough
// At this state, game will draw an error message
type SmallSizeState struct {
	Game   IGame
	Width  int
	Height int
}

// NewSmallSizeState returns a new SmallSizeState
func NewSmallSizeState(game IGame, width, height int) *SmallSizeState {
	return &SmallSizeState{game, width, height}
}

// OnResize checks if size is big enough
func (sss *SmallSizeState) OnResize(event *tcell.EventResize) {
	width, height := event.Size()
	if width >= sss.Game.MinWidth() && height >= sss.Game.MinHeight() {
		sss.Game.PopState()
	} else {
		sss.Width, sss.Height = width, height
	}
}

// OnKeyPress doesn't do anything
func (sss *SmallSizeState) OnKeyPress(event *tcell.EventKey) {}

// Draw draws the error message
func (sss *SmallSizeState) Draw() {
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
