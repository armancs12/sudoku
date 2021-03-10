package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/serhatscode/sudoku/board"
	"github.com/serhatscode/sudoku/ui"
)

// Game is an interface for the game
type Game interface {
	Start() error
	Exit()

	Board() board.Board
	SetBoard(board board.Board)

	MinWidth() int
	MinHeight() int

	State() State
	PushState(state State)
	PopState() State
}

type game struct {
	board  board.Board
	state  stateManager
	screen tcell.Screen

	minWidth, minHeight int
}

// NewGame returns a new game
func NewGame() (Game, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	game := &game{
		board:     board.New(board.Medium),
		screen:    screen,
		minWidth:  ui.BoardWidth,
		minHeight: ui.BoardHeight,
		state:     stateManager{},
	}

	game.state.Push(NewPlayState(game))

	return game, nil
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
	return game.state.Get()
}

func (game *game) PushState(state State) {
	game.state.Push(state)
}

func (game *game) PopState() State {
	return game.state.Pop()
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
