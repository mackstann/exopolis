package service

import (
	"github.com/mackstann/exopolis/city"
)

type CityService struct {
	city   *city.City
	zoning *city.ZoneMap
	jobs   *city.JobsLayer
	mapgen *city.MapGenerator
}

func NewCityService(c *city.City, zoning *city.ZoneMap, jobs *city.JobsLayer, mapgen *city.MapGenerator) *CityService {
	return &CityService{
		city:   c,
		zoning: zoning,
		jobs:   jobs,
		mapgen: mapgen,
	}
}

func (a *CityService) Step() {
	a.jobs.Step()
}

func (a *CityService) GenerateMap() {
	a.mapgen.Generate()
}
