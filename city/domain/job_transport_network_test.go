package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testCity4() City {
	city := NewCity(2)

	/* [road, farm]
	   [dirt, dirt]
	*/
	// XXX these changes don't affect the city if the network is set up before
	city[0][0].Typ = Road // u 0.9
	city[0][0].Road = &RoadCell{}
	city[0][1].Typ = Farm // temp .1
	city[0][1].Farm = &FarmCell{}
	return city
}

const (
	influxorWeight90 float64 = 1
	influxorWeight45 float64 = 1.0 / 4 // drops off with square of distance
)

func TestBasic(t *testing.T) {
	city := testCity4()

	jobTransport := NewJobTransportNetwork(city)
	jobTransport.Step()

	road := jobTransport.Grid[0][0]
	farm := jobTransport.Grid[0][1]

	// dirt has zero temp
	var expectedRoadTemp float64 = (farm.Temperature * road.conductivity) / (influxorWeight90 + influxorWeight90 + influxorWeight45)

	assert.InDelta(t, 0.04, expectedRoadTemp, 0.0001)

	// farm temp doesn't change
	assert.Equal(t, 0.1, farm.Temperature)
	// sink doesn't get hotter than source
	assert.GreaterOrEqual(t, farm.Temperature, road.Temperature)
	//
	assert.InDelta(t, expectedRoadTemp, road.Temperature, 0.0001)
}
