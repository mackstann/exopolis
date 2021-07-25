package domain

import (
	"log"

	cityDomain "github.com/mackstann/exopolis/city/domain"
)

type CityService interface {
	Get() *cityDomain.City
	Render() []string
	Step()
}

type MapGeneratorService interface {
	Generate(*cityDomain.City)
}

type InputPort interface {
	GetInputEventsNonBlocking() []InputEvent
}

type TerminalUIPort interface {
	TODORenderJustCity([]string)
}

type InputEvent int

const (
	QuitEvent     InputEvent = iota
	TODONoopEvent            = iota
)

type Game struct {
	city         CityService
	mapGenerator MapGeneratorService
	input        InputPort
	tui          TerminalUIPort
}

func NewGame(city CityService, mapGenerator MapGeneratorService, input InputPort, tui TerminalUIPort) *Game {
	return &Game{
		city:         city,
		mapGenerator: mapGenerator,
		input:        input,
		tui:          tui,
	}
}

func (g *Game) Run() {
	city := g.city.Get()
	g.mapGenerator.Generate(city)
	log.Println("game Run loop")

	// loop at time interval
	// non-blockingly check input chan
	// send msg to drawRequests
	for {
		for _, ev := range g.input.GetInputEventsNonBlocking() {
			if ev == QuitEvent {
				// TODO tell the terminal output to shutdown
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
