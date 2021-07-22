package main

import (
	cityDomain "github.com/mackstann/exopolis/city/domain"
	gameAdapters "github.com/mackstann/exopolis/game/adapters"
	gameDomain "github.com/mackstann/exopolis/game/domain"
)

func main() {
	city := cityDomain.NewCity(20)
	cityAdapter := gameAdapters.NewCityAdapter(city)
	mapGenerator := &mapgen{}
	terminal := gameAdapters.NewTerminalAdapter()
	game := gameDomain.NewGame(cityAdapter, mapGenerator, terminal)
	game.Run()
}

type mapgen struct{}

func (g *mapgen) Generate(city *cityDomain.City) {
	// XXX these changes don't affect the city if the network is set up before
	(*city)[0][0].Typ = cityDomain.House
	(*city)[0][0].House = &cityDomain.HouseCell{}
	(*city)[0][1].Typ = cityDomain.Road
	(*city)[0][1].Road = &cityDomain.RoadCell{}
	(*city)[0][2].Typ = cityDomain.PowerPlant
	(*city)[0][2].PowerPlant = &cityDomain.PowerPlantCell{}
	(*city)[0][3].Typ = cityDomain.Farm
	(*city)[0][3].Farm = &cityDomain.FarmCell{}
	(*city)[0][4].Typ = cityDomain.Farm
	(*city)[0][4].Farm = &cityDomain.FarmCell{}
	(*city)[1][2].Typ = cityDomain.Farm
	(*city)[1][2].Farm = &cityDomain.FarmCell{}
	(*city)[1][3].Typ = cityDomain.Farm
	(*city)[1][3].Farm = &cityDomain.FarmCell{}
	(*city)[1][4].Typ = cityDomain.Farm
	(*city)[1][4].Farm = &cityDomain.FarmCell{}
	(*city)[2][2].Typ = cityDomain.Farm
	(*city)[2][2].Farm = &cityDomain.FarmCell{}
	(*city)[2][3].Typ = cityDomain.Farm
	(*city)[2][3].Farm = &cityDomain.FarmCell{}
	(*city)[2][4].Typ = cityDomain.Farm
	(*city)[2][4].Farm = &cityDomain.FarmCell{}
}
