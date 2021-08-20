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
	zones := []string{
		"           ",
		"RRR        ",
		"RRR        ",
		"           ",
		"           ",
		"           ",
	}
	for y, row := range zones {
		for x, letter := range row {
			switch letter {
			case 'R':
				g.city.Zoning.SetZone(x, y, ResidentialZone)
			}
		}
	}

	// cells default to dirt
	city := []string{
		"PRRR   PFFF",
		"   R  FRFFF",
		"   R  FRFFF",
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
				g.city.Grid[y][x] = NewHouse()
			case 'R':
				g.city.Grid[y][x] = NewRoad()
			case 'F':
				g.city.Grid[y][x] = NewFarm()
			case 'P':
				g.city.Grid[y][x] = NewPowerPlant()
			}
		}
	}
}
