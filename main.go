package main

import (
	"time"

	"github.com/mackstann/exoplanet/city/domain"
	"github.com/mackstann/exoplanet/city/renderer"
)

func main() {
	city := domain.NewCity(5)

	// these changes don't affect the city if the network is set up before
	city[2][2].Typ = domain.House
	city[2][2].House = &domain.HouseCell{}
	city[3][2].Typ = domain.Road
	city[3][2].Road = &domain.RoadCell{}
	city[4][2].Typ = domain.Farm
	city[4][2].Farm = &domain.FarmCell{}

	jobTransport := domain.NewJobTransportNetwork(city)

	renderer.Render(city, jobTransport)

	realStep := 10 * time.Millisecond
	for {
		jobTransport.Step()
		renderer.Render(city, jobTransport)
		time.Sleep(realStep)
	}
}
