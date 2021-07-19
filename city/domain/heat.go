package domain

import (
	"fmt"
)

type HeatGrid struct {
	Grid []row
}

type cell struct {
	Temperature float64
	// conductivity controls heat flow into this cell. Heat flow out is unrestricted.
	conductivity float64
}

type row []cell

func NewHeatGrid(width int, height int) *HeatGrid {
	grid := make([]row, 0, height)
	for y := 0; y < height; y++ {
		grid = append(grid, make(row, width))
	}
	return &HeatGrid{
		Grid: grid,
	}
}

func (n HeatGrid) Step() {
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
			// influx, influxors := calcWeightedInflux(cells, weight)
			for _, c := range cells {
				if c.Temperature != -1 {
					fmt.Printf("90deg influxor\n")
					tote += c.Temperature
					influxors += 1
				}
			}
			for _, c := range diagCells {
				if c.Temperature != -1 {
					fmt.Printf("45deg influxor\n")
					tote += c.Temperature / 4
					influxors += 1.0 / 4
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

func (n HeatGrid) cellAt(x int, y int) cell {
	if x < 0 || y < 0 || x >= len(n.Grid) || y >= len(n.Grid[0]) {
		return cell{Temperature: -1, conductivity: -1}
	}
	return n.Grid[y][x]
}
