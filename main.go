package main

import (
	"log"
	"os"

	cityDomain "github.com/mackstann/exopolis/city/domain"
	renderer "github.com/mackstann/exopolis/city/renderer"
	gameAdapters "github.com/mackstann/exopolis/game/adapters"
	gameDomain "github.com/mackstann/exopolis/game/domain"
)

func main() {
	f, err := os.OpenFile("exopolis.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	city := cityDomain.NewCity(40)

	// problem: heat grid is now operating on an irrelevant grid
	network := cityDomain.NewJobTransportNetwork(city)

	cityAdapter := gameAdapters.NewCityAdapter(city, renderer.RenderTerminalTextBlock, network)
	mapGenerator := &mapgen{}
	terminal := gameAdapters.NewTerminalAdapter()
	game := gameDomain.NewGame(cityAdapter, mapGenerator, terminal, terminal)
	game.Run()
}

type mapgen struct{}

func (g *mapgen) Generate(city *cityDomain.City) {
	for y := 0; y < len(*city); y++ {
		for x := 0; x < len((*city)[0]); x++ {
			//dirt
			(*city)[y][x].JobConductivity = 0.1
		}
	}

	log.Printf("generating city")
	(*city)[0][0].Typ = cityDomain.House
	(*city)[0][0].House = &cityDomain.HouseCell{}
	(*city)[0][1].Typ = cityDomain.Road
	(*city)[0][1].Road = &cityDomain.RoadCell{}
	(*city)[0][1].JobConductivity = 0.9
	(*city)[0][2].Typ = cityDomain.PowerPlant
	(*city)[0][2].PowerPlant = &cityDomain.PowerPlantCell{}
	(*city)[0][2].JobTemperature = 1
	(*city)[0][3].Typ = cityDomain.Farm
	(*city)[0][3].Farm = &cityDomain.FarmCell{}
	(*city)[0][3].JobTemperature = 0.1
	(*city)[0][4].Typ = cityDomain.Farm
	(*city)[0][4].Farm = &cityDomain.FarmCell{}
	(*city)[0][4].JobTemperature = 0.1
	(*city)[1][2].Typ = cityDomain.Farm
	(*city)[1][2].Farm = &cityDomain.FarmCell{}
	(*city)[1][2].JobTemperature = 0.1
	(*city)[1][3].Typ = cityDomain.Farm
	(*city)[1][3].Farm = &cityDomain.FarmCell{}
	(*city)[1][3].JobTemperature = 0.1
	(*city)[1][4].Typ = cityDomain.Farm
	(*city)[1][4].Farm = &cityDomain.FarmCell{}
	(*city)[1][4].JobTemperature = 0.1
	(*city)[2][2].Typ = cityDomain.Farm
	(*city)[2][2].Farm = &cityDomain.FarmCell{}
	(*city)[2][2].JobTemperature = 0.1
	(*city)[2][3].Typ = cityDomain.Farm
	(*city)[2][3].Farm = &cityDomain.FarmCell{}
	(*city)[2][3].JobTemperature = 0.1
	(*city)[2][4].Typ = cityDomain.Farm
	(*city)[2][4].Farm = &cityDomain.FarmCell{}
	(*city)[2][4].JobTemperature = 0.1
}
