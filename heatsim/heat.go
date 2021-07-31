package heatsim

import (
	_ "log"
	"math"
)

// NewHeatGrid constructs a heat grid simulation.
//
// This is a simulation of heat flow in a 2-dimensional grid. The grid itself is not modeled here; it is only accessed
// by coordinates through TemperaturePort and ConductivityPort. The heat flow algorithm will automatically probe the
// boundaries of the grid without knowing its size; all it needs is for those two functions to return a float64 value (0
// to 1 inclusive) for a given (x, y), or return nil if the coordinates are out of bounds.
//
// efficiency (0 to 1 inclusive) is the proportion of heat that stays in the system as it flows from one cell to the
// next.  The remainder is lost as waste. Efficiency manifests as the distance heat can travel before it is lost to
// decay.
//
// The values returned by conductivityPort control the rate of heat flow into a given cell. Heat flow out is
// unrestricted. Conductivity manifests as the speed at which heat flows between cells.
//
// By tuning the efficiency of the grid and the temperature and conductivity of cells, different city-like phenomena can
// be simulated: traffic transmitting over a road system, electricity transmitting over power lines, crime and pollution
// radidating from their sources, etc.
//
// Surprising quirks:
// * Conductivity only affects heat flowing into a cell, not out.
// * While warm cells heat up cool cells, cool cells do not cool down warm cells. This enables heat to conduct over long
//   distances, almost more like electricity.
//
// Source of the dumbed down math used here:
// https://demonstrations.wolfram.com/ACellularAutomatonBasedHeatEquation/
//
// Inspired by SimCity (SNES).
func NewHeatGrid(efficiency float64, temp TemperaturePort, cond ConductivityPort) *HeatGrid {
	return &HeatGrid{
		efficiency:   efficiency,
		temperature:  temp,
		conductivity: cond,
	}
}

type HeatGrid struct {
	efficiency   float64
	temperature  TemperaturePort
	conductivity ConductivityPort
}

type TemperaturePort func(x int, y int) *float64
type ConductivityPort func(x int, y int) *float64

func (n HeatGrid) Step() {
	for y := 0; ; y++ {
		for x := 0; ; x++ {
			myTempPtr := n.temperature(x, y)
			if myTempPtr == nil {
				if x == 0 {
					return // end of rows
				}
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
