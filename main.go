package main

import (
	"github.com/mackstann/exoplanet/city/domain"
	"github.com/mackstann/exoplanet/city/renderer"
)

func main() {
	city := domain.NewCity(20)
	city[5][5].Typ = domain.House
	city[5][5].House = &domain.HouseCell{}
	city[6][5].Typ = domain.Farm
	city[6][5].Farm = &domain.FarmCell{}
	renderer.Render(city)
}
