package service

import (
	//"math/rand"

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
	//a.city.Step()
}

func (a *CityService) GenerateMap() {
	a.mapgen.Generate()
}

func (a *CityService) BuildResidential(x, y int) {
	a.city.Zoning.SetZone(x, y, city.ResidentialZone)
	a.city.Grid[y][x] = city.NewDirt()
}

/*
func occasionally() bool {
	return rand.Float64() < 0.001
}

func (c *City) Step() {
	for y := range c.Grid {
		for x := range c.Grid[y] {
			c.StepCell(x, y)
		}
	}
}

func (c *City) StepCell(x, y int) {
	row := c.Grid[y]
	cell := row[x]
	// move this logic into a rule...
	// needs to know about zoning AND cells...
	// it's a separate thing..? implemented by the city
	if c.Zoning.zoneAt(x, y) == ResidentialZone {
		if cell.Typ == Dirt &&
			cell.Resources.Jobs > 0.1 {
			if occasionally() {
				row[x] = NewHouse()
			}
		} else if cell.Typ == House &&
			cell.Resources.Jobs <= 0.1 {
			if occasionally() {
				row[x] = NewDirt()
			}
		}
	}
}
*/
