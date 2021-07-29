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
		"R      RFFF",
		"R      RFFF",
		"R      RFFF",
		"RRRRRRRRRRF",
	}

	log.Printf("generating city")

	// TODO: set .Jobs in cell type constructor

	// TODO: JobConductivity is remaining 0.1 for most cells on accident here
	// ??? is that true?

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
			}
		}
	}
}
