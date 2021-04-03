package ui

import "github.com/gdamore/tcell/v2"

func NewTCellClient() (Client, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	context := tcellContext{
		screen: screen,
		style:  tcell.StyleDefault,
	}

	client := tcellClient{
		context: context,
	}

	return &client, nil
}

var mapKeys = map[tcell.Key]string{
	tcell.KeyUp:    "arrow_up",
	tcell.KeyDown:  "arrow_down",
	tcell.KeyLeft:  "arrow_left",
	tcell.KeyRight: "arrow_right",
	tcell.KeyEnter: "enter",
	tcell.KeyESC:   "esc",
	tcell.KeyCtrlZ: "ctrl+z",
}

type tcellClient struct {
	context    tcellContext
	onResize   func(width, height int)
	onKeyPress func(key string)
}

func (tc *tcellClient) Start() error {
	err := tc.context.screen.Init()
	if err != nil {
		return err
	}

EventLoop:
	for {
		switch event := tc.waitForEvent().(type) {
		case *tcell.EventResize:
			tc.onResize(event.Size())
		case *tcell.EventKey:
			tc.onKeyPress(getGameKey(event))
		case nil:
			break EventLoop
		}
	}

	return nil
}

func (tc *tcellClient) Stop() {
	tc.context.screen.Fini()
}

func (tc *tcellClient) Size() (int, int) {
	return tc.context.screen.Size()
}

func (tc *tcellClient) OnResize(fn func(width, height int)) {
	tc.onResize = fn
}

func (tc *tcellClient) OnKeyPress(fn func(key string)) {
	tc.onKeyPress = fn
}

func (tc *tcellClient) Draw(x, y int, widget Widget) {
	widget.Draw(tc.Context(), x, y)
	tc.context.Show()
}

func (tc *tcellClient) DrawCenter(widget Widget) {
	width, height := tc.Size()
	x := (width - widget.Width()) / 2
	y := (height - widget.Height()) / 2

	tc.Draw(x, y, widget)
}

func (tc *tcellClient) DrawUpRightCorner(widget Widget) {
	width, _ := tc.Size()
	x := width - widget.Width()

	tc.Draw(x, 0, widget)
}

func (tc *tcellClient) Context() Context {
	return &tc.context
}

func (tc *tcellClient) waitForEvent() tcell.Event {
	return tc.context.screen.PollEvent()
}

type tcellContext struct {
	screen tcell.Screen
	style  tcell.Style
}

func (tc *tcellContext) StyleFG(color string) {
	tc.style.Foreground(tcell.GetColor(color))
}

func (tc *tcellContext) StyleBG(color string) {
	tc.style.Background(tcell.GetColor(color))
}

func (tc *tcellContext) SetContent(x, y int, char rune) {
	tc.screen.SetContent(x, y, char, nil, tc.style)
}

func (tc *tcellContext) Show() {
	tc.screen.Show()
}

func (tc *tcellContext) Clear() {
	tc.screen.Clear()
}

func getGameKey(event *tcell.EventKey) string {
	if event.Key() == tcell.KeyRune {
		return string(event.Rune())
	} else {
		return mapKeys[event.Key()]
	}
}
