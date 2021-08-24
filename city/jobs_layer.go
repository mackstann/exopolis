package city

import (
	"log"

	"github.com/mackstann/exopolis/heatsim"
)

type JobsLayer struct {
	city *City
	*heatsim.HeatGrid
	Grid JobsGrid
}

type JobsGrid [][]float64

const (
	dirtConductivity       float64 = 0.01
	roadConductivity               = 0.999
	powerPlantConductivity         = 0.9
	defaultConductivity            = 0.0
)

func NewJobsLayer(city *City) *JobsLayer {
	j := &JobsLayer{
		city: city,
		Grid: make([][]float64, 0, len(city.Grid)),
	}

	for range city.Grid {
		j.Grid = append(j.Grid, make([]float64, len(city.Grid[0])))
	}
	log.Printf("%v", j.Grid)

	temperature := func(x int, y int) (float64, bool) {
		if y < 0 || y >= len(j.Grid) || x < 0 || x >= len(j.Grid[0]) {
			log.Printf("temperature getter OOB 0")
			return 0, false
		}
		log.Printf("temperature getter %f", j.Grid[y][x])
		return j.Grid[y][x], true
	}
	setTemperature := func(x int, y int, val float64) {
		if y < 0 || y >= len(j.Grid) || x < 0 || x >= len(j.Grid[0]) {
			log.Panicf("setTemperature: out of bounds: (%d,%d)", x, y)
		}
		if city.Grid[y][x] == Farm {
			val = 0.1
		} else if city.Grid[y][x] == PowerPlant {
			val = 1
		}
		j.Grid[y][x] = val
	}
	getConductivity := func(x int, y int) (float64, bool) {
		if y < 0 || y >= len(j.Grid) || x < 0 || x >= len(j.Grid[0]) {
			return 0, false
		}

		switch city.Grid[y][x] {
		case Dirt, ResidentialZone, House:
			return dirtConductivity, true
		case Road:
			return roadConductivity, true
		case PowerPlant:
			return powerPlantConductivity, true
		default:
			return defaultConductivity, true
		}
	}
	j.HeatGrid = heatsim.NewHeatGrid(temperature, setTemperature, getConductivity)
	return j
}
