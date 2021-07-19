package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	influxorWeight90 float64 = 1
	influxorWeight45         = 1.0 / 4 // drops off with square of distance

	delta = 0.0001
)

func TestConduct90Degrees(t *testing.T) {
	heat := NewHeatGrid(2, 1)
	heater := &heat.Grid[0][0]
	conductor := &heat.Grid[0][1]
	heater.Temperature = 0.1
	conductor.conductivity = 0.9
	var expectedConductorTemp float64 = (heater.Temperature * conductor.conductivity) / influxorWeight90

	heat.Step()

	assert.InDelta(t, expectedConductorTemp, conductor.Temperature, delta, "conductor temperature is wrong")
}

func TestConduct45Degrees(t *testing.T) {
	heat := NewHeatGrid(2, 2)
	heater := &heat.Grid[0][0]
	conductor := &heat.Grid[1][1]
	heater.Temperature = 0.1
	conductor.conductivity = 0.9
	var expectedConductorTemp float64 = (heater.Temperature * conductor.conductivity / 4) / (influxorWeight90*2 + influxorWeight45)

	heat.Step()

	assert.InDelta(t, expectedConductorTemp, conductor.Temperature, delta, "conductor temperature is wrong")
}

func TestInsulatorNotHeated(t *testing.T) {
	heat := NewHeatGrid(2, 1)
	heater := &heat.Grid[0][0]
	insulator := &heat.Grid[0][1]
	heater.Temperature = 0.1

	heat.Step()

	assert.InDelta(t, 0, insulator.Temperature, delta, "insulator temperature is wrong")
}

func TestNonConductingHeaterNotHeated(t *testing.T) {
	heat := NewHeatGrid(2, 1)
	heater := &heat.Grid[0][0]
	heater.Temperature = 0.1

	heat.Step()

	assert.InDelta(t, 0.1, heater.Temperature, delta, "heater temp changed")
}
