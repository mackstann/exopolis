package domain

import (
	"math"
)

// HeatGrid is .. heat flows into each cell ...
type HeatGrid struct {
	Grid []row
	// efficiency is the proportion of heat that stays in the system as it flows from one cell to the next. The
	// remainder is lost as waste. Efficiency manifests as the distance heat can travel before it is lost to decay.
	efficiency float64
}

type cell struct {
	Temperature float64
	// conductivity controls the rate of heat flow into this cell. Heat flow out is unrestricted. Conductivity
	// manifests as the speed at which heat flows between cells.
	conductivity float64
}

type row []cell

func NewHeatGrid(width int, height int, efficiency float64) *HeatGrid {
	grid := make([]row, 0, height)
	for y := 0; y < height; y++ {
		grid = append(grid, make(row, width))
	}
	return &HeatGrid{
		Grid:       grid,
		efficiency: efficiency,
	}
}

func (n HeatGrid) Step() {
	for y := 0; y < len(n.Grid); y++ {
		for x := 0; x < len(n.Grid[y]); x++ {
			me := n.Grid[y][x]
			influx90, influxors90 := n.calcWeightedInflux(me.Temperature, n.neighbors90(x, y), 1)
			influx45, influxors45 := n.calcWeightedInflux(me.Temperature, n.neighbors45(x, y), 1.0/4)

			if influxors90 > 0 || influxors45 > 0 {
				ambientTemp := (influx90 + influx45) / (influxors90 + influxors45)

				tempDelta := ambientTemp - me.Temperature
				t := n.Grid[y][x].Temperature
				t += tempDelta * me.conductivity
				t = math.Min(1, t)
				t = math.Max(0, t)
				n.Grid[y][x].Temperature = t
			}
		}
	}
}

func (n HeatGrid) cellAt(x int, y int) cell {
	if x < 0 || y < 0 || y >= len(n.Grid) || x >= len(n.Grid[0]) {
		return cell{Temperature: -1, conductivity: -1}
	}
	return n.Grid[y][x]
}

func (n HeatGrid) calcWeightedInflux(myTemp float64, cells []cell, weight float64) (influx float64, influxors float64) {
	for _, c := range cells {
		if c.Temperature != -1 && c.Temperature > myTemp {
			influx += (c.Temperature * n.efficiency) * weight
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
