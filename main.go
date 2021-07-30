package main

import (
	"log"
	"os"
	"time"

	cityDomain "github.com/mackstann/exopolis/city"
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

	city := cityDomain.NewCity(20)

	// problem: heat grid is now operating on an irrelevant grid
	network := cityDomain.NewJobTransportNetwork(city)

	cityService := cityService.NewCityService(city, network, cityDomain.NewMapGenerator(city))
	terminal := gameAdapters.NewTerminalAdapter()
	renderer := gameAdapters.NewCityRenderer(city)

	cityService.GenerateMap()
	log.Println("game Run loop")

	for {
		for _, ev := range terminal.GetInputEventsNonBlocking() {
			if ev == game.QuitEvent {
				terminal.Shutdown()
				return
			}
		}
		for i := 0; i < 1; i++ {
			cityService.Step()
		}
		text := renderer.Render()
		terminal.UpdateCity(text)
		terminal.Redraw()
		// TODO: measure / compensate for frame processing time
		time.Sleep(time.Second / 60.0)
	}
}
