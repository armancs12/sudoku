package game

import (
	"github.com/serhatsdev/sudoku/game/board"
	"github.com/serhatsdev/sudoku/game/theme"
	"github.com/serhatsdev/sudoku/game/ui"
)

type Game interface {
	// Start starts the game
	Start() error
	// Exit exists the game
	Exit()

	// Board returns the current sudoku board
	Board() board.Board
	// SetBoard sets the current sudoku board
	SetBoard(board board.Board)

	// Theme returns the current theme
	Theme() theme.Theme
	// SetTheme sets current theme to the given theme
	SetTheme(theme theme.Theme)

	// MinWidth returns minimum terminal width
	// required to run the game
	MinWidth() int
	// MinHeight returns minimum terminal height
	// required to run the game
	MinHeight() int

	// Client returns the game client
	Client() ui.Client

	// State returns the current game state
	State() State
	// PushState sets the current game state
	// without replacing the previous one
	PushState(state State)
	// ChangeState sets the current state
	// to the given one
	ChangeState(state State)
	// PopState returns the current state and
	// sets the game state back to the previous state
	PopState() State
}

// NewGame returns a new game instance
func NewGame(client ui.Client) (Game, error) {
	themes, err := theme.GetThemes()
	if err != nil {
		return nil, err
	}

	game := &game{
		board:     board.New(board.Medium),
		client:    client,
		states:    []State{},
		minWidth:  ui.BoardWidth,
		minHeight: ui.BoardHeight,
		theme:     themes[0],
	}

	game.PushState(NewPlayState(game))

	return game, nil
}

type game struct {
	board  board.Board
	states []State
	client ui.Client
	theme  theme.Theme

	minWidth, minHeight int
}

func (game *game) Start() error {
	game.client.OnResize(func(width, height int) {
		game.client.Context().Clear()
		game.State().OnResize(width, height)
		game.State().Draw()
	})

	game.client.OnKeyPress(func(key string) {
		if key == "ctrl+z" {
			game.Exit()
		}

		game.client.Context().Clear()
		game.State().OnKeyPress(key)
		game.State().Draw()
	})

	err := game.client.Start()
	if err != nil {
		return err
	}

	return nil
}

func (game *game) Exit() {
	game.client.Stop()
}

func (game *game) Board() board.Board {
	return game.board
}

func (game *game) SetBoard(board board.Board) {
	game.board = board
}

func (game *game) Client() ui.Client {
	return game.client
}

func (game *game) Theme() theme.Theme {
	return game.theme
}

func (game *game) SetTheme(theme theme.Theme) {
	game.theme = theme
}

func (game *game) State() State {
	return game.states[len(game.states)-1]
}

func (game *game) PushState(state State) {
	game.states = append(game.states, state)
}

func (game *game) ChangeState(state State) {
	game.states[len(game.states)-1] = state
}

func (game *game) PopState() State {
	state := game.State()
	game.states = game.states[:len(game.states)-1]
	return state
}

func (game *game) MinWidth() int {
	return game.minWidth
}

func (game *game) MinHeight() int {
	return game.minHeight
}
