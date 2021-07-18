package domain

import (
	"time"
)

type JobTransportNetwork struct {
	city City
	Grid []row
}

func NewJobTransportNetwork(city City) *JobTransportNetwork {
	grid := make([]row, 0, len(city))
	for y := 0; y < len(city); y++ {
		grid = append(grid, make(row, len(city[y])))
	}

	for y := 0; y < len(city); y++ {
		for x := 0; x < len(city[y]); x++ {
			c := city[y][x]
			ourCell := cell{}
			switch c.Typ {
			case Dirt:
				ourCell.conductivity = 0.1
			case Road:
				ourCell.conductivity = 0.9
			case Farm:
				ourCell.Temperature = 1
			case PowerPlant:
				ourCell.Temperature = 10
			}
			grid[y][x] = ourCell
		}
	}

	return &JobTransportNetwork{
		city: city,
		Grid: grid,
	}
}

type cell struct {
	Temperature float64
	// conductivity controls heat flow into this cell. Heat flow out is unrestricted.
	conductivity float64
}

type row []cell

func (n JobTransportNetwork) Step(timeDelta time.Duration) {
	for y := 0; y < len(n.Grid); y++ {
		for x := 0; x < len(n.Grid[y]); x++ {
			cells := [4]cell{
				n.cellAt(x-1, y),
				n.cellAt(x+1, y),
				n.cellAt(x, y-1),
				n.cellAt(x, y+1),
			}
			diagCells := [4]cell{
				n.cellAt(x-1, y-1),
				n.cellAt(x-1, y+1),
				n.cellAt(x+1, y-1),
				n.cellAt(x+1, y+1),
			}

			me := n.Grid[y][x]
			var influx float64
			for _, c := range cells {
				temperatureDelta := c.Temperature - me.Temperature
				influx += temperatureDelta * me.conductivity
			}
			for _, c := range diagCells {
				temperatureDelta := c.Temperature - me.Temperature
				influx += temperatureDelta * me.conductivity / 4
			}

			// 1 each for right-angle neighbors, 1/4 each for diag neighbors
			avgInflux := influx / 5

			n.Grid[y][x].Temperature += avgInflux * float64(timeDelta)
		}
	}
}

func (n JobTransportNetwork) cellAt(x int, y int) cell {
	if x < 0 || y < 0 || x >= len(n.Grid) || y >= len(n.Grid[0]) {
		return cell{}
	}
	return n.Grid[y][x]
}

func seriesConductivity(c1 float64, c2 float64) float64 {
	if c1 == 0 || c2 == 0 {
		return 0
	}
	return (c1 * c2) / (c1 + c2)
}
