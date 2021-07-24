package adapters

import (
	cityDomain "github.com/mackstann/exopolis/city/domain"
)

type RenderFunc func(city cityDomain.City, n *cityDomain.JobTransportNetwork) []string

type CityAdapter struct {
	city *cityDomain.City
	// TODO belongs in city app layer!
	renderer RenderFunc
	network  *cityDomain.JobTransportNetwork
}

func NewCityAdapter(city *cityDomain.City, renderer RenderFunc, network *cityDomain.JobTransportNetwork) *CityAdapter {
	return &CityAdapter{
		city:     city,
		renderer: renderer,
		network:  network,
	}
}

func (a *CityAdapter) Get() *cityDomain.City {
	return a.city
}

func (a *CityAdapter) Render() []string {
	return a.renderer(*a.city, a.network)
}

func (a *CityAdapter) Step() {
	a.network.Step()
}
