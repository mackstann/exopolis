package domain

type CellType int

const (
	// Dirt, being the zero CellType value, is the default value of a newly allocated cell.
	Dirt CellType = 0

	Road      = 100
	PowerLine = 101

	House = 200
	Farm  = 201

	PowerPlant = 300
)

type Cell struct {
	Typ CellType

	Farm       *FarmCell
	House      *HouseCell
	PowerPlant *PowerPlantCell
	PowerLine  *PowerLineCell
	Road       *RoadCell
}

type FarmCell struct {
	productivity int
	accessTo     struct {
		water   int
		workers int
	}
}

type HouseCell struct {
	population int
	accessTo   struct {
		electricity int
		food        int
		work        int
	}
}

type PowerLineCell struct{}

type PowerPlantCell struct {
	production int
}

type RoadCell struct{}

type Row []Cell

type City []Row

func NewCity(size int) City {
	city := make(City, 0, size)
	for i := 0; i < size; i++ {
		city = append(city, make(Row, size))
	}
	return city
}
