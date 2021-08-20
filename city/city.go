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

func occasionally() bool {
	return rand.Float64() < 0.001
}

func (c *City) Step() {
	for y := range c.Grid {
		for x := range c.Grid[y] {
			c.StepCell(x, y)
		}
	}
}

func (c *City) StepCell(x, y int) {
	row := c.Grid[y]
	cell := row[x]
	// move this logic into a rule...
	// needs to know about zoning AND cells...
	// it's a separate thing..? implemented by the city
	if c.Zoning.zoneAt(x, y) == ResidentialZone {
		if cell.Typ == Dirt &&
			cell.Resources.Jobs > 0.1 {
			if occasionally() {
				row[x] = NewHouse()
			}
		} else if cell.Typ == House &&
			cell.Resources.Jobs <= 0.1 {
			if occasionally() {
				row[x] = NewDirt()
			}
		}
	}
}
