package domain

import (
	cityDomain "github.com/mackstann/exopolis/city/domain"
)

type CityService interface {
	Get() *cityDomain.City
}

type MapGeneratorService interface {
	Generate(*cityDomain.City)
}

type InputService interface {
	Events() chan InputEvent
}

type InputEvent int

const (
	QuitEvent InputEvent = iota
)

type Game struct {
	city         CityService
	mapGenerator MapGeneratorService
	input        InputService
}

func NewGame(city CityService, mapGenerator MapGeneratorService, input InputService) *Game {
	return &Game{
		city:         city,
		mapGenerator: mapGenerator,
		input:        input,
	}
}

func (g *Game) Run() {
	city := g.city.Get()
	g.mapGenerator.Generate(city)
	input := g.input.Events()
	for ev := range input {
		if ev == QuitEvent {
			// TODO tell the terminal output to shutdown
			return
		}
	}
}
