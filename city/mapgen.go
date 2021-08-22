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
		"PRRR   PFFF",
		"rrrR  FRFFF",
		"rrrR  FRFFF",
		"RRRR  FRFFF",
		"  RRRRRRRRF",
		"FFFFFFFFFFF",
	}

	log.Printf("generating city")

	// TODO: set .Jobs in cell type constructor

	for y, row := range city {
		for x, letter := range row {
			if y >= len(g.city.Grid) || x >= len(g.city.Grid[0]) {
				continue
			}
			switch letter {
			case 'H':
				g.city.Grid[y][x] = House
			case 'r':
				g.city.Grid[y][x] = ResidentialZone
			case 'R':
				g.city.Grid[y][x] = Road
			case 'F':
				g.city.Grid[y][x] = Farm
			case 'P':
				g.city.Grid[y][x] = PowerPlant
			}
		}
	}
}
