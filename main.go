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

// TODO: day/night cycle with coloration... or sun visualization

/* New approach... layers:
 * 1: build/input layer: what the player asked to build (includes zones)
 *   * currently called the city, but this is not accurate
 * 2..N: reaction layers -- electricity, traffic, houses etc.
 *   * these can have circular dependencies, but when they do, they should have a dampening effect so there aren't
 *     runaway effects.
 *
 * +growth -> +traffic -> -growth
 *
 * layers support a common interface so we can easily examine them in a sort of automated way
 *
 */

func main() {
	f, err := os.OpenFile("exopolis.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	const size = 20
	city := cityDomain.NewCity(size)
	jobs := cityDomain.NewJobsLayer(city)

	cityService := cityService.NewCityService(city, jobs, cityDomain.NewMapGenerator(city))
	terminal := gameAdapters.NewTerminalAdapter()
	renderer := gameAdapters.NewCityRenderer(city, jobs)
	g := game.NewGame(cityService, terminal)

	cityService.GenerateMap()
	log.Println("game Run loop")

	frameInterval := time.Duration(time.Second / 30.0)
	engineTickInterval := time.Duration(time.Second / 240.0)
	lastTick := time.Now()
	for {
		t := time.Now()
		for _, ev := range terminal.GetInputEventsNonBlocking() {
			if ev == game.QuitEvent {
				// TODO: move into game
				terminal.Shutdown()
				return
			} else {
				g.HandleInput(ev)
			}
		}
		tickDelta := t.Sub(lastTick)
		if tickDelta >= engineTickInterval {
			for tickDelta >= engineTickInterval {
				cityService.Step()
				tickDelta -= engineTickInterval
			}
			lastTick = time.Now()
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
		if frameInterval > duration {
			time.Sleep(frameInterval - duration)
		}
	}
}
