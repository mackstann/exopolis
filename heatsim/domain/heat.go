package domain

import (
	_ "log"
	"math"
)

// HeatGrid is .. heat flows into each cell ...
type HeatGrid struct {
	// efficiency is the proportion of heat that stays in the system as it flows from one cell to the next. The
	// remainder is lost as waste. Efficiency manifests as the distance heat can travel before it is lost to decay.
	efficiency float64

	temperature TemperaturePort
	// Conductivity controls the rate of heat flow into this cell. Heat flow out is unrestricted. Conductivity
	// manifests as the speed at which heat flows between cells.
	conductivity ConductivityPort
}

type TemperaturePort func(x int, y int) *float64
type ConductivityPort func(x int, y int) *float64

func NewHeatGrid(width int, height int, efficiency float64, temp TemperaturePort, cond ConductivityPort) *HeatGrid {
	return &HeatGrid{
		efficiency:   efficiency,
		temperature:  temp,
		conductivity: cond,
	}
}

func (n HeatGrid) Step() {
	for y := 0; ; y++ {
		if n.temperature(0, y) == nil {
			break // end of rows
		}
		for x := 0; ; x++ {
			//log.Printf("heat grid step: grid[%d][%d] is %v", y, x, n.Grid[y][x])
			myTempPtr := n.temperature(x, y)
			if myTempPtr == nil {
				break // end of line
			}
			myTemp := *myTempPtr
			x90, y90 := n.neighbors90(x, y)
			x45, y45 := n.neighbors45(x, y)
			influx90, influxors90 := n.calcWeightedInflux(myTemp, x90, y90, 1)
			influx45, influxors45 := n.calcWeightedInflux(myTemp, x45, y45, 1.0/4)

			if influxors90 > 0 || influxors45 > 0 {
				ambientTemp := (influx90 + influx45) / (influxors90 + influxors45)

				tempDelta := ambientTemp - myTemp
				myTemp += tempDelta * (*n.conductivity(x, y))
				myTemp = math.Min(1, myTemp)
				myTemp = math.Max(0, myTemp)
				*myTempPtr = myTemp
			}
		}
	}
}

func (n HeatGrid) calcWeightedInflux(myTemp float64, x []int, y []int, weight float64) (influx float64, influxors float64) {
	for i := range x {
		cx := x[i]
		cy := y[i]
		tempPtr := n.temperature(cx, cy)
		if tempPtr != nil && *tempPtr > myTemp {
			influx += (*tempPtr * n.efficiency) * weight
			influxors += 1 * weight
		}
	}
	return influx, influxors
}

func (n HeatGrid) neighbors90(x int, y int) ([]int, []int) {
	return []int{x - 1, x + 1, x, x},
		[]int{y, y, y - 1, y + 1}
}

func (n HeatGrid) neighbors45(x int, y int) ([]int, []int) {
	return []int{x - 1, x - 1, x + 1, x + 1},
		[]int{y - 1, y + 1, y - 1, y + 1}
}
