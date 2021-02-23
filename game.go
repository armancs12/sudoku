package sudoku

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/serhatscode/sudoku/ui"
)

// Game ...
type Game struct {
	board  Board
	state  StateManager
	screen tcell.Screen

	minWidth, minHeight int
}

// NewGame returns a new Game
func NewGame() (*Game, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	game := &Game{
		board:     NewBoard(40),
		screen:    screen,
		minWidth:  ui.BoardWidth,
		minHeight: ui.BoardHeight,
		state:     StateManager{},
	}

	game.state.Push(NewPlayState(game))

	return game, nil
}

// Start starts the game
func (game *Game) Start() error {
	err := ui.Init(game.screen)
	if err != nil {
		return err
	}

	for {
		switch event := game.screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.screen.Clear()
			game.state.Get().OnResize(event)
		case *tcell.EventKey:
			if event.Key() == tcell.KeyCtrlZ {
				game.cleanup()
			}
			game.state.Get().OnKeyPress(event)
		}
		game.state.Get().Draw()
	}
}

// cleanUp cleans the screen and exits
func (game *Game) cleanup() {
	game.screen.Fini()
	os.Exit(0)
}
