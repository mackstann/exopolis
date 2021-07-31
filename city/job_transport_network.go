package city

import (
	_ "log"

	"github.com/mackstann/exopolis/heatsim"
)

type JobTransportNetwork struct {
	city *City
	*heatsim.HeatGrid
}

const (
	dirtConductivity    float64 = 0.1
	roadConductivity            = 0.9
	defaultConductivity         = 0.0
)

func NewJobTransportNetwork(city *City) *JobTransportNetwork {
	temperature := func(x int, y int) *float64 {
		if y < 0 || y >= len(*city) || x < 0 || x >= len((*city)[0]) {
			return nil
		}
		// TODO: Try making temperature unmodifiable for some cells; see if it negates need for non-cooling hack
		return &(*city)[y][x].Resources.Jobs
	}
	conductivity := func(x int, y int) *float64 {
		if y < 0 || y >= len(*city) || x < 0 || x >= len((*city)[0]) {
			return nil
		}

		var c float64
		switch (*city)[y][x].Typ {
		case Dirt:
			c = dirtConductivity
		case Road:
			c = roadConductivity
		default:
			c = defaultConductivity
		}
		return &c
	}
	// TODO: pointer makes reads vs. writes mysterious
	// Use getter/setter. Conductivity only needs getter.
	// return secondary bool value in place of nil
	const efficiency = 0.9
	heat := heatsim.NewHeatGrid(efficiency, temperature, conductivity)
	return &JobTransportNetwork{
		city:     city,
		HeatGrid: heat,
	}
}
