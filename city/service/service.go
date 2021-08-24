package service

import (
	"log"
	"math/rand"

	"github.com/mackstann/exopolis/city"
)

type CityService struct {
	city   *city.City
	jobs   *city.JobsLayer
	mapgen *city.MapGenerator
}

func NewCityService(c *city.City, jobs *city.JobsLayer, mapgen *city.MapGenerator) *CityService {
	return &CityService{
		city:   c,
		jobs:   jobs,
		mapgen: mapgen,
	}
}

func (a *CityService) Step() {
	a.jobs.Step()

	for y := range a.city.Grid {
		for x := range a.city.Grid[y] {
			a.StepCell(x, y)
		}
	}
}

func (a *CityService) GenerateMap() {
	a.mapgen.Generate()
}

func (a *CityService) BuildResidential(x, y int) {
	a.city.Grid[y][x] = city.ResidentialZone
}

func (a *CityService) BuildRoad(x, y int) {
	a.city.Grid[y][x] = city.Road
}

func (a *CityService) BuildPowerPlant(x, y int) {
	a.city.Grid[y][x] = city.PowerPlant
}

func occasionally() bool {
	return rand.Float64() < 0.001
}

func (a *CityService) StepCell(x, y int) {
	row := a.city.Grid[y]
	cell := row[x]
	// move this logic into a rule...
	// needs to know about zoning AND cells...
	// it's a separate thing..? implemented by the city
	if cell == city.ResidentialZone &&
		a.jobs.Grid[y][x] > 0.1 {
		log.Printf("grow")
		if occasionally() {
			row[x] = city.House
		}
	} else if cell == city.House &&
		a.jobs.Grid[y][x] <= 0.1 {
		if occasionally() {
			log.Printf("ungrow")
			row[x] = city.ResidentialZone
		}
	}
}
