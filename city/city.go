package city

type Row []Cell

type City struct {
	Grid   []Row
	Zoning *ZoneMap
}

func NewCity(size int, zoning *ZoneMap) *City {
	grid := make([]Row, 0, size)
	for i := 0; i < size; i++ {
		grid = append(grid, make(Row, size))
	}
	return &City{
		Grid:   grid,
		Zoning: zoning,
	}
}
