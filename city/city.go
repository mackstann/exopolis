package city

import (
	"math/rand"
)

type Row []Cell

type City struct {
	Grid   []Row
	Zoning *ZoneMap
}

func NewCity(size int, zoning *ZoneMap) *City {
	grid := make([]Row, 0, size)
	for i := 0; i < size; i++ {
		grid = append(grid, make(Row, size))
	}
	return &City{
		Grid:   grid,
		Zoning: zoning,
	}
}

// TODO: call Step()
// TODO: add zoning map generator
func (c *City) Step() {
	for y, row := range c.Grid {
		for x, cell := range c.Grid[y] {
			if c.Zoning.zoneAt(x, y) == ResidentialZone &&
				cell.Typ == Dirt &&
				cell.Resources.Jobs > 0.1 {
				r := rand.Float64()
				if r < 0.001 {
					row[x].Typ = House
				}
			}
		}
	}
}
