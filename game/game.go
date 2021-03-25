package game

import (
	"os"

	"github.com/gdamore/tcell/v2"
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

	// MinWidth returns minimum terminal width
	// required to run the game
	MinWidth() int
	// MinHeight returns minimum terminal height
	// required to run the game
	MinHeight() int

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
func NewGame() (Game, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	game := &game{
		board:        board.New(board.Medium),
		screen:       screen,
		minWidth:     ui.BoardWidth,
		minHeight:    ui.BoardHeight,
		stateManager: stateManager{},
	}

	game.stateManager.Push(NewPlayState(game))

	return game, nil
}

type game struct {
	board        board.Board
	stateManager stateManager
	screen       tcell.Screen

	minWidth, minHeight int
}

func (game *game) Start() error {
	err := ui.Init(game.screen)
	if err != nil {
		return err
	}

	for {
		switch event := game.screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.screen.Clear()
			game.State().OnResize(event)
		case *tcell.EventKey:
			if event.Key() == tcell.KeyCtrlZ {
				game.Exit()
			}
			game.State().OnKeyPress(event)
		}
		game.State().Draw()
	}
}

func (game *game) Exit() {
	game.screen.Fini()
	os.Exit(0)
}

func (game *game) Board() board.Board {
	return game.board
}

func (game *game) SetBoard(board board.Board) {
	game.board = board
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
