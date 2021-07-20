package domain

type JobTransportNetwork struct {
	city City
	*HeatGrid
}

func NewJobTransportNetwork(city City) *JobTransportNetwork {
	const efficiency = 0.9
	heat := NewHeatGrid(len(city[0]), len(city), efficiency)
	for y := 0; y < len(city); y++ {
		for x := 0; x < len(city[y]); x++ {
			c := city[y][x]
			heatCell := &heat.Grid[y][x]
			switch c.Typ {
			case Dirt:
				heatCell.conductivity = 0.1
			case Road:
				heatCell.conductivity = 0.9
			case Farm:
				heatCell.Temperature = 0.1
			case PowerPlant:
				heatCell.Temperature = 1
			}
		}
	}

	return &JobTransportNetwork{
		city:     city,
		HeatGrid: heat,
	}
}
