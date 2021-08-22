package city

type ZoneType int

const (
	NoZone          ZoneType = 0
	ResidentialZone          = 1
	CommercialZone           = 2
	IndustrialZone           = 3
)

type ZoneMap [][]ZoneType

func NewZoneMap(size int) *ZoneMap {
	zmap := make(ZoneMap, 0, size)
	for i := 0; i < size; i++ {
		zmap = append(zmap, make([]ZoneType, size))
	}
	return &zmap
}

func (z *ZoneMap) ZoneAt(x int, y int) ZoneType {
	return (*z)[y][x]
}

func (z *ZoneMap) SetZone(x int, y int, zt ZoneType) {
	(*z)[y][x] = zt
}
