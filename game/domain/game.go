package domain

import (
	"log"

	"github.com/mackstann/exopolis/city"
)

type CityService interface {
	GenerateMap()
	Render() []string
	Step()
}

type MapGeneratorService interface {
	Generate(*city.City)
}

type InputPort interface {
	GetInputEventsNonBlocking() []InputEvent
}

type TerminalUIPort interface {
	TODORenderJustCity([]string)
	Shutdown()
}

type InputEvent int

const (
	QuitEvent     InputEvent = iota
	TODONoopEvent            = iota
)

type Game struct {
	city  CityService
	input InputPort
	tui   TerminalUIPort
}

func NewGame(city CityService, input InputPort, tui TerminalUIPort) *Game {
	return &Game{
		city:  city,
		input: input,
		tui:   tui,
	}
}

func (g *Game) Run() {
	g.city.GenerateMap()
	log.Println("game Run loop")

	for {
		for _, ev := range g.input.GetInputEventsNonBlocking() {
			if ev == QuitEvent {
				g.tui.Shutdown()
				return
			}
		}
		for i := 0; i < 1; i++ {
			g.city.Step()
		}
		text := g.city.Render()
		g.tui.TODORenderJustCity(text)
	}
}
