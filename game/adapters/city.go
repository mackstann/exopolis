package adapters

import (
	"github.com/mackstann/exopolis/city"
)

type RenderFunc func(city.City, *city.JobTransportNetwork) []string

type CityAdapter struct {
	city *city.City
	// TODO belongs in city app layer!
	renderer RenderFunc
	network  *city.JobTransportNetwork
}

func NewCityAdapter(c *city.City, renderer RenderFunc, network *city.JobTransportNetwork) *CityAdapter {
	return &CityAdapter{
		city:     c,
		renderer: renderer,
		network:  network,
	}
}

func (a *CityAdapter) Get() *city.City {
	return a.city
}

func (a *CityAdapter) Render() []string {
	return a.renderer(*a.city, a.network)
}

func (a *CityAdapter) Step() {
	a.network.Step()
}
