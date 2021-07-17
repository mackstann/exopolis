package domain

import (
	"time"
)

type jobTransportNetwork struct {
	city City
	grid []row
}

func NewJobTransportNetwork(city City) {
	grid = make([]row, 0, len(n.city))
	for y := 0; y < len(n.city); y++ {
		grid = append(grid, make(row, len(n.city[y])))
	}

	for y := 0; y < len(n.city); y++ {
		for x := 0; x < len(n.city[y]); x++ {
			c := n.city[y][x]
			ourCell := cell{}
			switch c.Typ {
			case Dirt:
				ourCell.conductivity = 0.1
			case Road:
				ourCell.conductivity = 0.9
			case Farm:
				ourCell.heat = 1
			case PowerPlant:
				ourCell.heat = 10
			}
			grid[y][x] = ourCell
		}
	}

	return &jobTransportNetwork{
		city: City,
		grid: grid,
	}
}

type cell struct {
	heat         float64
	conductivity float64
}

type row []cell

func (n jobTransportNetwork) step(timeDelta time.Time) {
	for y := 0; y < len(n.grid); y++ {
		for x := 0; x < len(n.grid[y]); x++ {
			influx := (n.influx(x-1, y) +
				n.influx(x+1, y) +
				n.influx(x, y-1) +
				n.influx(x, y+1) +
				n.influx(x-1, y-1)/4 +
				n.influx(x-1, y+1)/4 +
				n.influx(x+1, y-1)/4 +
				n.influx(x+1, y+1)/4)
			influx /= 5
			//n.grid[y][x].heat = // combine their heat and ours, depending on our conductivity
		}
	}
}

func (n jobTransportNetwork) influx(x int, y int) {
	if x < 0 || y < 0 || x >= len(n.grid) || y >= len(n.grid[0]) {
		return 0
	}
	cell := n.grid[y][x]
	return cell.heat * cell.conductivity

}
