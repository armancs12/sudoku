package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/serhatscode/sudoku/board"
	"github.com/serhatscode/sudoku/ui"
)

// IGame is an interface for the game
type IGame interface {
	Start() error
	Exit()

	Board() board.IBoard
	SetBoard(board board.IBoard)

	MinWidth() int
	MinHeight() int

	State() State
	PushState(state State)
	PopState() State
}

// Game ...
type Game struct {
	board  board.IBoard
	state  StateManager
	screen tcell.Screen

	minWidth, minHeight int
}

// NewGame returns a new Game
func NewGame() (IGame, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	game := &Game{
		board:     board.New(board.Medium),
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

// Exit cleans the screen and exits the game
func (game *Game) Exit() {
	game.screen.Fini()
	os.Exit(0)
}

// Board returns the sudoku board
func (game *Game) Board() board.IBoard {
	return game.board
}

// SetBoard sets the sudoku board
func (game *Game) SetBoard(board board.IBoard) {
	game.board = board
}

// State returns the current game state
func (game *Game) State() State {
	return game.state.Get()
}

// PushState pushes new state to on top of state stack
func (game *Game) PushState(state State) {
	game.state.Push(state)
}

// PopState pops the state from state stack
func (game *Game) PopState() State {
	return game.state.Pop()
}

// MinWidth returns minimum width
// for the game to be able to run properly
func (game *Game) MinWidth() int {
	return game.minWidth
}

// MinHeight returns minimum height
// for the game to be able to run properly
func (game *Game) MinHeight() int {
	return game.minHeight
}

// StateManager is an array of states
type StateManager []State

// Get returns the current game state
// which is the last element of the array
func (sm *StateManager) Get() State {
	if len(*sm) > 0 {
		return (*sm)[len(*sm)-1]
	}
	return nil
}

// Push appends the given state
func (sm *StateManager) Push(state State) {
	*sm = append(*sm, state)
}

// Pop removes the last element
func (sm *StateManager) Pop() State {
	state := sm.Get()
	if len(*sm) > 0 {
		*sm = (*sm)[:len(*sm)-1]
	}
	return state
}
