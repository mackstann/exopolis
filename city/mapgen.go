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
	for y := 0; y < len(*g.city); y++ {
		for x := 0; x < len((*g.city)[0]); x++ {
			//dirt
			(*g.city)[y][x].JobConductivity = 0.1
		}
	}

	log.Printf("generating city")
	(*g.city)[0][0].Typ = House
	(*g.city)[0][1].Typ = Road
	(*g.city)[0][1].JobConductivity = 0.9
	(*g.city)[0][2].Typ = PowerPlant
	(*g.city)[0][2].JobTemperature = 1
	(*g.city)[0][3].Typ = Farm
	(*g.city)[0][3].JobTemperature = 0.1
	(*g.city)[0][4].Typ = Farm
	(*g.city)[0][4].JobTemperature = 0.1
	(*g.city)[1][2].Typ = Farm
	(*g.city)[1][2].JobTemperature = 0.1
	(*g.city)[1][3].Typ = Farm
	(*g.city)[1][3].JobTemperature = 0.1
	(*g.city)[1][4].Typ = Farm
	(*g.city)[1][4].JobTemperature = 0.1
	(*g.city)[2][2].Typ = Farm
	(*g.city)[2][2].JobTemperature = 0.1
	(*g.city)[2][3].Typ = Farm
	(*g.city)[2][3].JobTemperature = 0.1
	(*g.city)[2][4].Typ = Farm
	(*g.city)[2][4].JobTemperature = 0.1
}
