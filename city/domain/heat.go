package domain

import (
	"fmt"
	"math"
)

// HeatGrid is .. heat flows into each cell ...
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
			influx90, influxors90 := calcWeightedInflux(n.neighbors90(x, y), 1)
			influx45, influxors45 := calcWeightedInflux(n.neighbors45(x, y), 1.0/4)

			ambientTemp := (influx90 + influx45) / (influxors90 + influxors45)

			me := n.Grid[y][x]
			tempDelta := ambientTemp - me.Temperature
			if y == 0 && x == 1 {
				// the problem is that the road is pulling in the coldness of other things around it
				fmt.Printf("i am %v, ambient is %v, temp to add is %v via conductivity %v\n", me.Temperature, ambientTemp, tempDelta, me.conductivity)
			}
			t := n.Grid[y][x].Temperature
			t += tempDelta * me.conductivity
			t = math.Min(1, t)
			t = math.Max(0, t)
			n.Grid[y][x].Temperature = t
		}
	}
}

func (n HeatGrid) cellAt(x int, y int) cell {
	if x < 0 || y < 0 || y >= len(n.Grid) || x >= len(n.Grid[0]) {
		return cell{Temperature: -1, conductivity: -1}
	}
	return n.Grid[y][x]
}

func calcWeightedInflux(cells []cell, weight float64) (influx float64, influxors float64) {
	for _, c := range cells {
		if c.Temperature != -1 {
			influx += c.Temperature * weight
			influxors += 1 * weight
		}
	}
	return influx, influxors
}

func (n HeatGrid) neighbors90(x int, y int) []cell {
	return []cell{
		n.cellAt(x-1, y),
		n.cellAt(x+1, y),
		n.cellAt(x, y-1),
		n.cellAt(x, y+1),
	}
}

func (n HeatGrid) neighbors45(x int, y int) []cell {
	return []cell{
		n.cellAt(x-1, y-1),
		n.cellAt(x-1, y+1),
		n.cellAt(x+1, y-1),
		n.cellAt(x+1, y+1),
	}
}
