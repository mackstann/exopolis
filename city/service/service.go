package service

import (
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
	a.city.Step()
}

func (a *CityService) GenerateMap() {
	a.mapgen.Generate()
}
