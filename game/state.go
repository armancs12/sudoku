package game

import (
	"github.com/serhatsdev/sudoku/game/board"
)

// State is an interface for game states
// The current game state will handle:
//
// - What to do when game resized
// - What to do when key pressed
// - What to draw
type State interface {
	// OnResize will be invoked when terminal resized
	OnResize(width, height int)
	// OnKeyPress will be invoked when key pressed
	OnKeyPress(key string)
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
		Options: []menuOption{
			{"Resume", func() {
				game.PopState()
			}},
			{"New Game", func() {
				game.SetBoard(board.New(board.Medium))
				game.PopState()
			}},
			{"Themes", func() {
				themes, err := LoadThemes()
				if err != nil {
					return
				}

				game.PushState(NewThemesMenuState(game, themes))
			}},
			{"Exit", func() {
				game.Exit()

			}},
		},
	}
}

func NewThemesMenuState(game Game, themes []Theme) State {
	themeOptions := []menuOption{}
	for _, theme := range themes {
		theme := theme

		themeOptions = append(themeOptions, menuOption{
			title: theme.Name,
			function: func() {
				game.SetTheme(theme)
				game.PopState()
			},
		})
	}

	return &menuState{
		Game:    game,
		Options: themeOptions,
	}
}
