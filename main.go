package main

import (
	"log"
	"os"

	"github.com/mackstann/exopolis/city"
	cityAdapters "github.com/mackstann/exopolis/city/adapters"
	cityService "github.com/mackstann/exopolis/city/service"
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

	c := city.NewCity(40)

	// problem: heat grid is now operating on an irrelevant grid
	network := city.NewJobTransportNetwork(c)

	cityService := cityService.NewCityService(c, cityAdapters.Render, network)
	mapGenerator := &mapgen{}
	terminal := gameAdapters.NewTerminalAdapter()
	game := gameDomain.NewGame(cityService, mapGenerator, terminal, terminal)
	game.Run()
}

type mapgen struct{}

func (g *mapgen) Generate(c *city.City) {
	for y := 0; y < len(*c); y++ {
		for x := 0; x < len((*c)[0]); x++ {
			//dirt
			(*c)[y][x].JobConductivity = 0.1
		}
	}

	log.Printf("generating city")
	(*c)[0][0].Typ = city.House
	(*c)[0][0].House = &city.HouseCell{}
	(*c)[0][1].Typ = city.Road
	(*c)[0][1].Road = &city.RoadCell{}
	(*c)[0][1].JobConductivity = 0.9
	(*c)[0][2].Typ = city.PowerPlant
	(*c)[0][2].PowerPlant = &city.PowerPlantCell{}
	(*c)[0][2].JobTemperature = 1
	(*c)[0][3].Typ = city.Farm
	(*c)[0][3].Farm = &city.FarmCell{}
	(*c)[0][3].JobTemperature = 0.1
	(*c)[0][4].Typ = city.Farm
	(*c)[0][4].Farm = &city.FarmCell{}
	(*c)[0][4].JobTemperature = 0.1
	(*c)[1][2].Typ = city.Farm
	(*c)[1][2].Farm = &city.FarmCell{}
	(*c)[1][2].JobTemperature = 0.1
	(*c)[1][3].Typ = city.Farm
	(*c)[1][3].Farm = &city.FarmCell{}
	(*c)[1][3].JobTemperature = 0.1
	(*c)[1][4].Typ = city.Farm
	(*c)[1][4].Farm = &city.FarmCell{}
	(*c)[1][4].JobTemperature = 0.1
	(*c)[2][2].Typ = city.Farm
	(*c)[2][2].Farm = &city.FarmCell{}
	(*c)[2][2].JobTemperature = 0.1
	(*c)[2][3].Typ = city.Farm
	(*c)[2][3].Farm = &city.FarmCell{}
	(*c)[2][3].JobTemperature = 0.1
	(*c)[2][4].Typ = city.Farm
	(*c)[2][4].Farm = &city.FarmCell{}
	(*c)[2][4].JobTemperature = 0.1
}
