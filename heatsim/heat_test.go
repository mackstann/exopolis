package heatsim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	influxorWeight90 float64 = 1
	influxorWeight45         = 1.0 / 4 // drops off with square of distance

	delta                = 0.0001
	conductionEfficiency = 0.9
)

type cell struct {
	x            int
	y            int
	temperature  float64
	conductivity float64
}

func makeTemperaturePort(width int, height int, cells []*cell) TemperaturePort {
	return func(x int, y int) *float64 {
		if x < 0 || x >= width || y < 0 || y >= height {
			return nil
		}
		for _, c := range cells {
			if c.x == x && c.y == y {
				return &(*c).temperature
			}
		}
		return new(float64)
	}
}

func makeConductivityPort(width int, height int, cells []*cell) ConductivityPort {
	return func(x int, y int) (float64, bool) {
		if x < 0 || x >= width || y < 0 || y >= height {
			return 0, false
		}
		for _, c := range cells {
			if c.x == x && c.y == y {
				return (*c).conductivity, true
			}
		}
		return 0, true
	}
}

func TestConduct90Degrees(t *testing.T) {
	heater := cell{x: 0, y: 0, temperature: 0.1}
	conductor := cell{x: 1, y: 0, conductivity: 0.9}
	cells := []*cell{&heater, &conductor}
	heat := NewHeatGrid(conductionEfficiency, makeTemperaturePort(2, 1, cells), makeConductivityPort(2, 1, cells))

	heat.Step()

	var expectedConductorTemp float64 = (heater.temperature * conductor.conductivity * conductionEfficiency) / influxorWeight90
	assert.InDelta(t, expectedConductorTemp, conductor.temperature, delta, "conductor temperature is wrong")
}

func TestConduct45Degrees(t *testing.T) {
	heater := cell{x: 0, y: 0, temperature: 0.1}
	conductor := cell{x: 1, y: 1, conductivity: 0.9}
	cells := []*cell{&heater, &conductor}
	heat := NewHeatGrid(conductionEfficiency, makeTemperaturePort(2, 2, cells), makeConductivityPort(2, 2, cells))

	heat.Step()

	var expectedConductorTemp float64 = (heater.temperature * conductor.conductivity * conductionEfficiency / 4) / influxorWeight45
	assert.InDelta(t, expectedConductorTemp, conductor.temperature, delta, "conductor temperature is wrong")
}

func TestInsulatorNotHeated(t *testing.T) {
	heater := cell{x: 0, y: 0, temperature: 0.1}
	insulator := cell{x: 1, y: 0}
	cells := []*cell{&heater, &insulator}
	heat := NewHeatGrid(conductionEfficiency, makeTemperaturePort(2, 1, cells), makeConductivityPort(2, 1, cells))

	heat.Step()

	assert.InDelta(t, 0, insulator.temperature, delta, "insulator temperature is wrong")
}

func TestNonConductingHeaterNotHeated(t *testing.T) {
	heater := cell{x: 0, y: 0, temperature: 0.1}
	cells := []*cell{&heater}
	heat := NewHeatGrid(conductionEfficiency, makeTemperaturePort(2, 1, cells), makeConductivityPort(2, 1, cells))

	heat.Step()

	assert.InDelta(t, 0.1, heater.temperature, delta, "heater temp changed")
}

func TestCoolerCellDoesNotHeatMe(t *testing.T) {
	warm := cell{x: 0, y: 0, temperature: 0.2}
	cool := cell{x: 1, y: 0, temperature: 0.1}
	cells := []*cell{&warm, &cool}
	heat := NewHeatGrid(conductionEfficiency, makeTemperaturePort(2, 1, cells), makeConductivityPort(2, 1, cells))

	heat.Step()

	assert.InDelta(t, 0.2, warm.temperature, delta, "warmer cell got cooled")
}
