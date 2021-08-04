package adapters

import (
	"fmt"
	"log"

	"github.com/mackstann/exopolis/city"

	"github.com/muesli/termenv"
)

type CityRenderer struct {
	city *city.City
}

func NewCityRenderer(c *city.City) *CityRenderer {
	return &CityRenderer{
		city: c,
	}
}

func (r *CityRenderer) Render() []string {
	log.Println("render it")
	return textualize(*r.city)
}

func textualize(c city.City) []string {
	rows := make([]string, 0, len(c.Grid))
	for y, row := range c.Grid {
		rowOutput := ""
		for x, cell := range row {
			temp255 := int(c.Grid[y][x].Resources.Jobs * 255.0)
			intensity := fmt.Sprintf("%02x", temp255)
			chr := "."
			color := ""
			switch cell.Typ {
			case city.House:
				chr = "■"
				color = intensity + "0000"
			case city.Farm:
				chr = "▤"
				color = "00" + intensity + "00"
			case city.Road:
				chr = "▪"
				color = intensity + intensity + intensity
			case city.PowerPlant:
				chr = "p"
				color = intensity + intensity + "00"
			case city.Dirt:
				chr = "░"
				color = "0000" + intensity
			}
			p := termenv.ColorProfile()
			if intensity != "" && chr != " " && len(color) == 6 {
				rowOutput += termenv.String(chr).Foreground(p.Color("#" + color)).String()
			} else {
				rowOutput += chr
			}
		}
		rows = append(rows, rowOutput)
	}
	return rows
}
