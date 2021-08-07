package game

import (
	"log"
)

type InputEvent int

const (
	QuitEvent InputEvent = iota
	CursorUp
	CursorDown
	CursorLeft
	CursorRight
)

type Game struct {
	view    ViewPort
	cursorX int
	cursorY int
}

type ViewPort interface {
	MoveCursor(x, y int)
}

func NewGame(view ViewPort) *Game {
	return &Game{view: view}
}

func (g *Game) HandleInput(ev InputEvent) {
	switch ev {
	case CursorUp:
		if g.cursorY > 0 {
			g.cursorY--
			g.view.MoveCursor(g.cursorX, g.cursorY)
		}
	case CursorDown:
		if g.cursorY < 19 { // TODO
			g.cursorY++
			g.view.MoveCursor(g.cursorX, g.cursorY)
		}
	case CursorLeft:
		log.Printf("CursorLeft")
		if g.cursorX > 0 {
			g.cursorX--
			g.view.MoveCursor(g.cursorX, g.cursorY)
		}
	case CursorRight:
		if g.cursorX < 19 { // TODO
			g.cursorX++
			g.view.MoveCursor(g.cursorX, g.cursorY)
		}
	}
	// TODO: a game renderer that composites this on top...
	// except... can't use background as cursor if we composite like that
}
