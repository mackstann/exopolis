package city

import (
	"log"
)

func NewMapGenerator(c *City) *MapGenerator {
	return &MapGenerator{city: c}
}

type MapGenerator struct {
	city *City
}

func (g *MapGenerator) Generate() {
	// cells default to dirt

	log.Printf("generating city")
	// TODO: JobConductivity is remaining 0.1 for most cells on accident here
	(*g.city)[0][0].Typ = House
	(*g.city)[0][1].Typ = Road
	(*g.city)[0][2].Typ = PowerPlant
	(*g.city)[0][2].Resources.Jobs = 1
	(*g.city)[0][3].Typ = Farm
	(*g.city)[0][3].Resources.Jobs = 0.1
	(*g.city)[0][4].Typ = Farm
	(*g.city)[0][4].Resources.Jobs = 0.1
	(*g.city)[1][2].Typ = Farm
	(*g.city)[1][2].Resources.Jobs = 0.1
	(*g.city)[1][3].Typ = Farm
	(*g.city)[1][3].Resources.Jobs = 0.1
	(*g.city)[1][4].Typ = Farm
	(*g.city)[1][4].Resources.Jobs = 0.1
	(*g.city)[2][2].Typ = Farm
	(*g.city)[2][2].Resources.Jobs = 0.1
	(*g.city)[2][3].Typ = Farm
	(*g.city)[2][3].Resources.Jobs = 0.1
	(*g.city)[2][4].Typ = Farm
	(*g.city)[2][4].Resources.Jobs = 0.1
}
