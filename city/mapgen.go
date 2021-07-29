package city

import (
	"log"
)

func NewMapGenerator(c *City) *MapGenerator {
	return &MapGenerator{city: c}
}

type MapGenerator struct {
	city *City
}

func (g *MapGenerator) Generate() {
	// cells default to dirt
	city := []string{
		"H      PFFF",
		"R     FRFFF",
		"R     FRFFF",
		"R     FRFFF",
		"RRRRRRRRRRF",
		"FFFFFFFFFFF",
	}

	log.Printf("generating city")

	// TODO: set .Jobs in cell type constructor

	for y, row := range city {
		for x, letter := range row {
			switch letter {
			case 'H':
				(*g.city)[y][x].Typ = House
			case 'R':
				(*g.city)[y][x].Typ = Road
			case 'F':
				(*g.city)[y][x].Typ = Farm
				(*g.city)[y][x].Resources.Jobs = 0.1
			case 'P':
				(*g.city)[y][x].Typ = PowerPlant
				(*g.city)[y][x].Resources.Jobs = 1
				/* conundrum:
				if we have 12 farms = 1.2 jobs,
				and we have 1 power plant = 1 job,
				then why does the power plant transmit so much more job power?
				even when the farms are all along the roads

				maybe because temperature does not scale linearly?
				how do we map between it and linear counting of things?
				*/
			}
		}
	}
}
