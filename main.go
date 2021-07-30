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
		t := time.Now()
		for _, ev := range terminal.GetInputEventsNonBlocking() {
			if ev == game.QuitEvent {
				terminal.Shutdown()
				return
			}
		}
		for i := 0; i < 1; i++ {
			// TODO: give the game engine its own clock, independent of rendering
			cityService.Step()
		}
		text := renderer.Render()
		// TODO: for other UI elements, composite multiple layers of text.
		terminal.UpdateCity(text)
		terminal.Redraw()
		tEnd := time.Now()

		duration := tEnd.Sub(t)
		desiredDuration := time.Duration(time.Second / 30.0)
		if desiredDuration > duration {
			time.Sleep(desiredDuration - duration)
		}
	}
}
