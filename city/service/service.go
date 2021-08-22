package service

import (
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
	a.city.Zoning.SetZone(x, y, city.ResidentialZone)
	a.city.Grid[y][x] = city.NewDirt()
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
	if a.city.Zoning.ZoneAt(x, y) == city.ResidentialZone {
		if cell.Typ == city.Dirt &&
			a.jobs.Grid[y][x] > 0.1 {
			if occasionally() {
				row[x] = city.NewHouse()
			}
		} else if cell.Typ == city.House &&
			a.jobs.Grid[y][x] <= 0.1 {
			if occasionally() {
				row[x] = city.NewDirt()
			}
		}
	}
}
