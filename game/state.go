package game

import (
	"github.com/gdamore/tcell/v2"
	"github.com/serhatscode/sudoku/board"
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

// NewPlayState returns a new play state
func NewPlayState(game Game) State {
	return &playState{
		Game: game,
		// center of the board
		Pos: board.Point2{X: 4, Y: 4},
	}
}

// NewSmallSizeState returns a new small size state
func NewSmallSizeState(game Game, width, height int) State {
	return &smallSizeState{game, width, height}
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
