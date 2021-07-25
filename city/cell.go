package city

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

// Networks (implemented as graph):
// * foodTransport (has resistance)
// * jobTransport (has resistance)
// * electrical (no resistance)
//
// Fields (implemented as heat diffusion)
// * Air pollution
// * Noise
//
// Transpo demand/traffic...
//
// day/night cycle with coloration... or sun visualization

// level is determined by field strengths
// e.g. formula: house level is (electricalBool x foodTransport x jobTransport x (1-pollution) x (1-noise))

type Cell struct {
	Typ CellType

	JobConductivity float64
	JobTemperature  float64

	Farm       *FarmCell
	House      *HouseCell
	PowerPlant *PowerPlantCell
	PowerLine  *PowerLineCell
	Road       *RoadCell
}

type FarmCell struct{}

type HouseCell struct{}

type PowerLineCell struct{}

type PowerPlantCell struct{}

type RoadCell struct{}
