package city

type CellType int

const (
	// Dirt, being the zero CellType value, is the default value of a newly allocated cell.
	Dirt CellType = 0

	Road      = 100
	PowerLine = 101

	House = 200
	// TODO: farm doesn't make sense in a city
	Farm = 201

	PowerPlant = 300
)

type Resources struct {
	Jobs       float64
	JobsSource bool
}

type Cell struct {
	Typ       CellType
	Resources Resources
}

func NewDirt() Cell {
	return Cell{Typ: Dirt}
}

func NewHouse() Cell {
	return Cell{Typ: House}
}

func NewRoad() Cell {
	return Cell{Typ: Road}
}

func NewFarm() Cell {
	return Cell{
		Typ: Farm,
		Resources: Resources{
			Jobs:       0.1,
			JobsSource: true,
		},
	}
}

func NewPowerPlant() Cell {
	return Cell{
		Typ: PowerPlant,
		Resources: Resources{
			Jobs:       1,
			JobsSource: true,
		},
	}
}
