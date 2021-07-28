package game

import (
	"log"
)

type CityService interface {
	GenerateMap()
	Step()
}

type CityRendererPort interface {
	Render() []string
}

type InputPort interface {
	GetInputEventsNonBlocking() []InputEvent
}

type TerminalUIPort interface {
	UpdateCity([]string)
	Redraw()
	Shutdown()
}

type InputEvent int

const (
	QuitEvent InputEvent = iota
)

type Game struct {
	city         CityService
	input        InputPort
	tui          TerminalUIPort
	cityRenderer CityRendererPort
}

func NewGame(city CityService, input InputPort, tui TerminalUIPort, cityRenderer CityRendererPort) *Game {
	return &Game{
		city:         city,
		input:        input,
		tui:          tui,
		cityRenderer: cityRenderer,
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
		/* this is weird...
		should the city service know what datatype render returns?
		should the game service know it either?
		interactor -> presenter -> view
		then the interactor and view are ignorant of each others' formats
		but let's skip the view?

		we can know about city stuff, but only through its service API
		and make sure cross-feature deps are a DAG
		game -> city -> heatsim
		*/
		text := g.cityRenderer.Render()
		g.tui.UpdateCity(text)
		g.tui.Redraw()
	}
}
