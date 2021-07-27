package main

import (
	"log"
	"os"

	"github.com/mackstann/exopolis/city"
	cityService "github.com/mackstann/exopolis/city/service"
	"github.com/mackstann/exopolis/game"
	gameAdapters "github.com/mackstann/exopolis/game/adapters"
)

func main() {
	f, err := os.OpenFile("exopolis.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	c := city.NewCity(40)

	// problem: heat grid is now operating on an irrelevant grid
	network := city.NewJobTransportNetwork(c)

	cityService := cityService.NewCityService(c, network, city.NewMapGenerator(c))
	terminal := gameAdapters.NewTerminalAdapter()
	renderer := gameAdapters.NewCityRenderer(c)
	game := game.NewGame(cityService, terminal, terminal, renderer)
	game.Run()
}
