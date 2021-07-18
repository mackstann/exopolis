package domain

import (
	"fmt"
	"time"
)

type JobTransportNetwork struct {
	city City
	Grid []row
}

const (
	influxorWeight90 float64 = 1.0
	// because the heat transmission force reduces with the square of the distance
	influxorWeight45 = 1.0 / 4
)

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
				ourCell.Temperature = 0.1
			case PowerPlant:
				ourCell.Temperature = 1
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

			fmt.Printf("-------\n")
			me := n.Grid[y][x]
			var tote float64
			var influxors float64
			for _, c := range cells {
				if c.Temperature != -1 {
					fmt.Printf("90deg influxor\n")
					tote += c.Temperature
					influxors += influxorWeight90
				}
			}
			for _, c := range diagCells {
				if c.Temperature != -1 {
					fmt.Printf("45deg influxor\n")
					tote += c.Temperature / 4
					influxors += influxorWeight45
				}
			}

			fmt.Printf("influxors %f\n", influxors)
			avg := tote / influxors
			delta := avg - me.Temperature
			// DELTA IS 0.5 but i thought it should be 0.1??????
			if delta > 0 {
				fmt.Printf("metemp %v delta %v cond %v\n", me.Temperature, delta, me.conductivity)
				n.Grid[y][x].Temperature += delta * me.conductivity
				if n.Grid[y][x].Temperature > 1 {
					n.Grid[y][x].Temperature = 1
				}
				if n.Grid[y][x].Temperature < 0 {
					n.Grid[y][x].Temperature = 0
				}
			}
			fmt.Printf("-------\n")
		}
	}
}

func (n JobTransportNetwork) cellAt(x int, y int) cell {
	if x < 0 || y < 0 || x >= len(n.Grid) || y >= len(n.Grid[0]) {
		return cell{Temperature: -1, conductivity: -1}
	}
	return n.Grid[y][x]
}

func seriesConductivity(c1 float64, c2 float64) float64 {
	if c1 == 0 || c2 == 0 {
		return 0
	}
	return (c1 * c2) / (c1 + c2)
}
