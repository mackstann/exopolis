package city

import (
	"github.com/mackstann/exopolis/heatsim"
)

type JobTransportNetwork struct {
	city *City
	*heatsim.HeatGrid
}

func NewJobTransportNetwork(city *City) *JobTransportNetwork {
	const efficiency = 0.9
	temperature := func(x int, y int) *float64 {
		if y < 0 || y >= len(*city) || x < 0 || x >= len((*city)[0]) {
			return nil
		}
		return &(*city)[y][x].JobTemperature
	}
	conductivity := func(x int, y int) *float64 {
		if y < 0 || y >= len(*city) || x < 0 || x >= len((*city)[0]) {
			return nil
		}
		return &(*city)[y][x].JobConductivity
	}
	heat := heatsim.NewHeatGrid(len((*city)[0]), len((*city)), efficiency, temperature, conductivity)
	return &JobTransportNetwork{
		city:     city,
		HeatGrid: heat,
	}
}
