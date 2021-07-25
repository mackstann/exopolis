package adapters

import (
	"github.com/mackstann/exopolis/city"
)

type RenderFunc func(city.City, *city.JobTransportNetwork) []string

type CityService struct {
	city     *city.City
	renderer RenderFunc
	network  *city.JobTransportNetwork
	mapgen   *city.MapGenerator
}

func NewCityService(c *city.City, renderer RenderFunc, network *city.JobTransportNetwork, mapgen *city.MapGenerator) *CityService {
	return &CityService{
		city:     c,
		renderer: renderer,
		network:  network,
		mapgen:   mapgen,
	}
}

func (a *CityService) Render() []string {
	return a.renderer(*a.city, a.network)
}

func (a *CityService) Step() {
	a.network.Step()
}

func (a *CityService) GenerateMap() {
	a.mapgen.Generate()
}
