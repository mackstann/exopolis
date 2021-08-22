package city

type Cell int

const (
	// Dirt, being the zero Cell value, is the default value of a newly allocated cell.
	Dirt Cell = 0

	Road      = 100
	PowerLine = 101

	ResidentialZone = 200
	House           = 201

	// TODO: farm doesn't make sense in a city
	Farm = 300

	PowerPlant = 400
)
