package city

type Cell int

const (
	// Dirt, being the zero Cell value, is the default value of a newly allocated cell.
	Dirt Cell = 0

	Road      = 100
	PowerLine = 101

	House = 200
	// TODO: farm doesn't make sense in a city
	Farm = 201

	PowerPlant = 300
)
