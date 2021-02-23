package sudoku

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
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

// PlayState is the state of gameplay
type PlayState struct {
	game       *Game
	posX, posY int
}

// NewPlayState returns a new PlayState
func NewPlayState(game *Game) *PlayState {
	return &PlayState{
		game: game,
		// center of the board
		posX: 4,
		posY: 4,
	}
}

// OnResize checks if size is big enough
func (ps *PlayState) OnResize(event *tcell.EventResize) {
	width, height := event.Size()
	if width < ps.game.minWidth || height < ps.game.minHeight {
		ps.game.state.Push(NewSmallSizeState(ps.game, width, height))
	}
}

// OnKeyPress changes the board indicator position
func (ps *PlayState) OnKeyPress(event *tcell.EventKey) {
	key := event.Key()

	if key == tcell.KeyESC {
		ps.game.state.Push(NewMenuState(ps.game))
		return
	}

	if key == tcell.KeyUp && ps.posY > 0 {
		ps.posY--
	} else if key == tcell.KeyDown && ps.posY < 8 {
		ps.posY++
	} else if key == tcell.KeyLeft && ps.posX > 0 {
		ps.posX--
	} else if key == tcell.KeyRight && ps.posX < 8 {
		ps.posX++
	} else {
		char := event.Rune()
		if char == 'e' || char == 'E' {
			ps.game.board[ps.posY][ps.posX] = 0
		} else {
			num := int(char - '0')
			if num > 0 && num < 10 {
				ps.game.board[ps.posY][ps.posX] = num
			}
		}
	}
}

// Draw draws the board
func (ps *PlayState) Draw() {
	ui.DrawCenter(&ui.BoardWidget{
		Board:       ps.game.board,
		CursorX:     ps.posX,
		CursorY:     ps.posY,
		CursorColor: tcell.ColorRed,
		NumberColor: tcell.ColorBlack,
		BorderColor: tcell.ColorBlack,
		Background:  tcell.ColorWhite,
	})
}

// MenuState is the state of game menu
type MenuState struct {
	game      *Game
	pos       int
	options   []string
	functions []func()
}

// NewMenuState returns a new MenuState
func NewMenuState(game *Game) *MenuState {

	return &MenuState{
		game: game,
		pos:  0,
		options: []string{
			"Resume",
			"New Game",
			"Exit",
		},
		functions: []func(){
			func() {
				game.state.Pop()
			},
			func() {
				game.board = NewBoard(40)
				game.state.Pop()
			},
			func() {
				game.cleanup()
			},
		},
	}
}

// OnResize checks if size is big enough
func (ms *MenuState) OnResize(event *tcell.EventResize) {
	width, height := event.Size()
	if width < ms.game.minWidth || height < ms.game.minHeight {
		ms.game.state.Push(NewSmallSizeState(ms.game, width, height))
	}
}

// OnKeyPress changes the menu indicator position
func (ms *MenuState) OnKeyPress(event *tcell.EventKey) {
	key := event.Key()
	if key == tcell.KeyESC {
		ms.game.state.Pop()
		return
	}

	if key == tcell.KeyUp {
		ms.pos = (len(ms.options) + ms.pos - 1) % len(ms.options)
	} else if key == tcell.KeyDown {
		ms.pos = (ms.pos + 1) % len(ms.options)
	} else if key == tcell.KeyEnter {
		ms.functions[ms.pos]()
	}
}

// Draw draws the game menu
func (ms *MenuState) Draw() {
	ui.DrawCenter(&ui.BoxWidget{
		Child: &ui.MenuWidget{
			Options:          ms.options,
			CursorIndex:      ms.pos,
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
	game   *Game
	width  int
	height int
}

// NewSmallSizeState returns a new SmallSizeState
func NewSmallSizeState(game *Game, width, height int) *SmallSizeState {
	return &SmallSizeState{game, width, height}
}

// OnResize checks if size is big enough
func (sss *SmallSizeState) OnResize(event *tcell.EventResize) {
	width, height := event.Size()
	if width >= sss.game.minWidth && height >= sss.game.minHeight {
		sss.game.state.Pop()
	} else {
		sss.width, sss.height = width, height
	}
}

// OnKeyPress doesn't do anything
func (sss *SmallSizeState) OnKeyPress(event *tcell.EventKey) {}

// Draw draws the error message
func (sss *SmallSizeState) Draw() {
	message := fmt.Sprintf("Please resize to\n at least %dx%d",
		sss.game.minWidth, sss.game.minHeight)
	current := fmt.Sprintf("%dx%d", sss.width, sss.height)

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
