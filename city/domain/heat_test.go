package domain

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

func TestConduct90Degrees(t *testing.T) {
	heat := NewHeatGrid(2, 1, conductionEfficiency)
	heater := &heat.Grid[0][0]
	conductor := &heat.Grid[0][1]
	heater.Temperature = 0.1
	conductor.conductivity = 0.9
	var expectedConductorTemp float64 = (heater.Temperature * conductor.conductivity * conductionEfficiency) / influxorWeight90

	heat.Step()

	assert.InDelta(t, expectedConductorTemp, conductor.Temperature, delta, "conductor temperature is wrong")
}

func TestConduct45Degrees(t *testing.T) {
	heat := NewHeatGrid(2, 2, conductionEfficiency)
	heater := &heat.Grid[0][0]
	conductor := &heat.Grid[1][1]
	heater.Temperature = 0.1
	conductor.conductivity = 0.9
	var expectedConductorTemp float64 = (heater.Temperature * conductor.conductivity * conductionEfficiency / 4) / influxorWeight45

	heat.Step()

	assert.InDelta(t, expectedConductorTemp, conductor.Temperature, delta, "conductor temperature is wrong")
}

func TestInsulatorNotHeated(t *testing.T) {
	heat := NewHeatGrid(2, 1, conductionEfficiency)
	heater := &heat.Grid[0][0]
	insulator := &heat.Grid[0][1]
	heater.Temperature = 0.1

	heat.Step()

	assert.InDelta(t, 0, insulator.Temperature, delta, "insulator temperature is wrong")
}

func TestNonConductingHeaterNotHeated(t *testing.T) {
	heat := NewHeatGrid(2, 1, conductionEfficiency)
	heater := &heat.Grid[0][0]
	heater.Temperature = 0.1

	heat.Step()

	assert.InDelta(t, 0.1, heater.Temperature, delta, "heater temp changed")
}

func TestCoolerCellDoesNotHeatMe(t *testing.T) {
	heat := NewHeatGrid(2, 1, conductionEfficiency)
	warm := &heat.Grid[0][0]
	warm.Temperature = 0.2
	cool := &heat.Grid[0][1]
	cool.Temperature = 0.1

	heat.Step()

	assert.InDelta(t, 0.2, warm.Temperature, delta, "warmer cell got cooled")
}
