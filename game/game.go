package game

import (
	_ "log"
)

type InputEvent int

const (
	QuitEvent InputEvent = iota
)

type Game struct {
}

func NewGame() *Game {
	return &Game{}
}
