package main

import (
	"github.com/mackstann/exoplanet/city/domain"
	"github.com/mackstann/exoplanet/city/renderer"
)

func main() {
	city := domain.NewCity(10)
	jobTransport := domain.NewJobTransportNetwork(city)

	city[5][5].Typ = domain.House
	city[5][5].House = &domain.HouseCell{}
	city[6][5].Typ = domain.Road
	city[6][5].Farm = &domain.RoadCell{}
	city[7][5].Typ = domain.Farm
	city[7][5].Farm = &domain.FarmCell{}
	renderer.Render(city)

	jobTransport.step()

	renderer.Render(city)
}
