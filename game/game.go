package game

import (
	"github.com/serhatscode/sudoku/board"
	"github.com/serhatscode/sudoku/ui"
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
	Theme() Theme
	// SetTheme sets current theme to the given theme
	SetTheme(theme Theme)

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
	// PopState returns the current state and
	// sets the game state back to the previous state
	PopState() State
}

// NewGame returns a new game instance
func NewGame(client ui.Client) (Game, error) {
	game := &game{
		board:        board.New(board.Medium),
		client:       client,
		minWidth:     ui.BoardWidth,
		minHeight:    ui.BoardHeight,
		stateManager: stateManager{},
		theme:        DefaultTheme(),
	}

	game.stateManager.Push(NewPlayState(game))

	return game, nil
}

type game struct {
	board        board.Board
	stateManager stateManager
	client       ui.Client
	theme        Theme

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

func (game *game) Theme() Theme {
	return game.theme
}

func (game *game) SetTheme(theme Theme) {
	game.theme = theme
}

func (game *game) State() State {
	return game.stateManager.Get()
}

func (game *game) PushState(state State) {
	game.stateManager.Push(state)
}

func (game *game) PopState() State {
	return game.stateManager.Pop()
}

func (game *game) MinWidth() int {
	return game.minWidth
}

func (game *game) MinHeight() int {
	return game.minHeight
}

// StateManager is an array of states
type stateManager []State

// Get returns the current game state
// which is the last element of the array
func (sm *stateManager) Get() State {
	if len(*sm) > 0 {
		return (*sm)[len(*sm)-1]
	}
	return nil
}

// Push appends the given state
func (sm *stateManager) Push(state State) {
	*sm = append(*sm, state)
}

// Pop removes the last element
func (sm *stateManager) Pop() State {
	state := sm.Get()
	if len(*sm) > 0 {
		*sm = (*sm)[:len(*sm)-1]
	}
	return state
}
