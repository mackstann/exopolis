package heatsim

import (
	"log"
	"math"
)

// NewHeatGrid constructs a heat grid simulation.
//
// This is a simulation of heat flow in a 2-dimensional grid, with some unrealistic modifications to make it more useful
// in a game. The grid itself is not modeled here; it is only accessed by coordinates through the temperature and
// conductivity interfaces. The heat flow algorithm will automatically probe the boundaries of the grid without knowing
// its size; all it needs is for those two functions to return a float64 value (0 to 1 inclusive) for a given (x, y),
// along with a bool indicating whether the coordinates are in bounds.
//
// The values returned by conductivityPort control the rate of heat flow into a given cell. Heat flow out is
// unrestricted. Conductivity manifests as the speed at which heat flows between cells.
//
// By tuning the temperature and conductivity of cells, different city-like phenomena can be simulated: traffic
// transmitting over a road system, electricity transmitting over power lines, crime and pollution radidating from their
// sources, etc.
//
// Inspired by SimCity (SNES).
func NewHeatGrid(temp TemperatureGetter, setTemp TemperatureSetter, cond ConductivityGetter) *HeatGrid {
	return &HeatGrid{
		temperature:    temp,
		setTemperature: setTemp,
		conductivity:   cond,
	}
}

type HeatGrid struct {
	temperature    TemperatureGetter
	setTemperature TemperatureSetter
	conductivity   ConductivityGetter
}

type TemperatureGetter func(x int, y int) (float64, bool)
type TemperatureSetter func(x int, y int, val float64)
type ConductivityGetter func(x int, y int) (float64, bool)

func (n HeatGrid) Step() {
	for y := 0; ; y++ {
		for x := 0; ; x++ {
			myTemp, hasTemp := n.temperature(x, y)
			if !hasTemp {
				if x == 0 {
					return // end of rows
				}
				break // end of line
			}
			conductivity, hasConductivity := n.conductivity(x, y)
			if !hasConductivity {
				log.Panicf("already checked cell (%d,%d); should always have conductivity", x, y)
			}
			x90, y90 := n.neighbors90(x, y)
			x45, y45 := n.neighbors45(x, y)
			influx90, influxors90 := n.calcWeightedInflux(myTemp, conductivity, x90, y90, 1)
			influx45, influxors45 := n.calcWeightedInflux(myTemp, conductivity, x45, y45, 1.0/4)

			if influxors90 > 0 || influxors45 > 0 {
				log.Printf("myTemp is %v", myTemp)
				log.Printf("influx %v, influxors %v", (influx90 + influx45), (influxors90 + influxors45))
				log.Printf("add weighted influx %v", (influx90+influx45)/(influxors90+influxors45))
				myTemp += (influx90 + influx45) / (influxors90 + influxors45)
				myTemp = math.Min(1, myTemp)
				myTemp = math.Max(0, myTemp)
				n.setTemperature(x, y, myTemp)
			}
		}
	}
}

func (n HeatGrid) calcWeightedInflux(myTemp float64, myConductivity float64, x []int, y []int, weight float64) (influx float64, influxors float64) {
	for i := range x {
		cx := x[i]
		cy := y[i]
		temp, hasTemp := n.temperature(cx, cy)
		cond, hasCond := n.conductivity(cx, cy)
		if temp != 0 {
			log.Printf("found temp %f BUT hasTemp %v, cond %f, hasCond %v", temp, hasTemp, cond, hasCond)
		}
		if hasTemp && hasCond && myConductivity > 0 && cond > 0 {
			log.Printf("influx calc %f %f %v %v", temp, myTemp, seriesConductivity(myConductivity, cond), weight)
			limitingConductivity := math.Max(myConductivity, cond)
			goalTemp := temp * limitingConductivity
			deltaTemp := goalTemp - myTemp
			influx += deltaTemp * weight * seriesConductivity(myConductivity, cond)
			influxors += weight
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

func seriesConductivity(c1 float64, c2 float64) float64 {
	if c1 == 0 || c2 == 0 {
		return 0
	}
	return (c1 * c2) / (c1 + c2)
}
