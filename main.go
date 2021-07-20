package main

import (
	"time"

	"github.com/mackstann/exoplanet/city/domain"
	"github.com/mackstann/exoplanet/city/renderer"
)

func main() {
	city := domain.NewCity(20)

	// these changes don't affect the city if the network is set up before
	city[0][0].Typ = domain.House
	city[0][0].House = &domain.HouseCell{}
	city[0][1].Typ = domain.Road
	city[0][1].Road = &domain.RoadCell{}
	city[0][2].Typ = domain.PowerPlant
	city[0][2].PowerPlant = &domain.PowerPlantCell{}
	city[0][3].Typ = domain.Farm
	city[0][3].Farm = &domain.FarmCell{}
	city[0][4].Typ = domain.Farm
	city[0][4].Farm = &domain.FarmCell{}
	city[1][2].Typ = domain.Farm
	city[1][2].Farm = &domain.FarmCell{}
	city[1][3].Typ = domain.Farm
	city[1][3].Farm = &domain.FarmCell{}
	city[1][4].Typ = domain.Farm
	city[1][4].Farm = &domain.FarmCell{}
	city[2][2].Typ = domain.Farm
	city[2][2].Farm = &domain.FarmCell{}
	city[2][3].Typ = domain.Farm
	city[2][3].Farm = &domain.FarmCell{}
	city[2][4].Typ = domain.Farm
	city[2][4].Farm = &domain.FarmCell{}

	jobTransport := domain.NewJobTransportNetwork(city)

	renderer.Render(city, jobTransport)

	realStep := 10 * time.Millisecond
	for {
		jobTransport.Step()
		renderer.Render(city, jobTransport)
		time.Sleep(realStep)
	}
}
