package adapters

import (
	"fmt"
	_ "log"

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

func (r *CityRenderer) Render() [][]string {
	c := *r.city
	rows := make([][]string, 0, len(c.Grid))
	for y, row := range c.Grid {
		rowOutput := make([]string, 0, len(c.Grid[y]))
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
				rowOutput = append(rowOutput, termenv.String(chr).Foreground(p.Color("#"+color)).String())
			} else {
				rowOutput = append(rowOutput, chr)
			}
		}
		rows = append(rows, rowOutput)
	}
	return rows
}
