package adapters

import (
	"github.com/mackstann/exopolis/city"
)

type CityService struct {
	city    *city.City
	network *city.JobTransportNetwork
	mapgen  *city.MapGenerator
}

func NewCityService(c *city.City, network *city.JobTransportNetwork, mapgen *city.MapGenerator) *CityService {
	return &CityService{
		city:    c,
		network: network,
		mapgen:  mapgen,
	}
}

func (a *CityService) Step() {
	a.network.Step()
}

func (a *CityService) GenerateMap() {
	a.mapgen.Generate()
}
