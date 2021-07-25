package adapters

import (
	"github.com/mackstann/exopolis/city"
)

type RenderFunc func(city.City, *city.JobTransportNetwork) []string

type CityService struct {
	city     *city.City
	renderer RenderFunc
	network  *city.JobTransportNetwork
}

func NewCityService(c *city.City, renderer RenderFunc, network *city.JobTransportNetwork) *CityService {
	return &CityService{
		city:     c,
		renderer: renderer,
		network:  network,
	}
}

// TODO this shouldn't be exposed
func (a *CityService) Get() *city.City {
	return a.city
}

func (a *CityService) Render() []string {
	return a.renderer(*a.city, a.network)
}

func (a *CityService) Step() {
	a.network.Step()
}
