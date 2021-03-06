package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

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
	logEnabled := flag.Bool("log", false, "")
	flag.Parse()
	if logEnabled != nil && *logEnabled {
		f, err := os.OpenFile("exopolis.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("error opening log file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)
	} else {
		log.SetOutput(ioutil.Discard)
	}

	const size = 20

	cityService := cityService.NewCityService(size)

	// TODO game factory func
	terminal := gameAdapters.NewTerminalAdapter()
	g := game.NewGame(cityService, terminal)

	// TODO game should call when starting
	cityService.GenerateMap()

	log.Println("game Run loop")

	frameInterval := time.Duration(time.Second / 30.0)
	engineTickInterval := time.Duration(time.Second / 240.0)
	lastTick := time.Now()
	for {
		t := time.Now()
		for _, ev := range terminal.GetInputEventsNonBlocking() {
			g.HandleInput(ev)
			if g.Done() {
				return
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
			cityService.Step()
		}
		text := cityService.Render()
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
