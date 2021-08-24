package adapters

import (
	"log"
	"math/rand"

	"github.com/mackstann/exopolis/city"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/termenv"
)

const (
	// wong color palette
	black  = "#000000"
	orange = "#E69F00"
	aqua   = "#56B4E9"
	green  = "#009E73"
	yellow = "#F0E442"
	blue   = "#0072B2"
	red    = "#D55E00"
	pink   = "#CC79A7"
	// supplemental
	grey = "#aaaaaa"
)

type CityRenderer struct {
	city *city.City
	jobs *city.JobsLayer
}

func NewCityRenderer(c *city.City, jobs *city.JobsLayer) *CityRenderer {
	return &CityRenderer{
		city: c,
		jobs: jobs,
	}
}

func (r *CityRenderer) Render() [][]string {
	c := *r.city
	rows := make([][]string, 0, len(c.Grid))
	p := termenv.ColorProfile()
	for y, row := range c.Grid {
		rowOutput := make([]string, 0, len(c.Grid[y]))
		for x, cell := range row {
			log.Printf("render %d %d", x, y)
			log.Printf("r.jobs %v", r.jobs)
			chr := "."
			color := ""
			switch cell {
			case city.ResidentialZone:
				chr = "▀"
				color = blue
			case city.House:
				chr = "▄"
				color = pink
			case city.Farm:
				chr = "▤"
				color = green
			case city.Road:
				chr = "▪"
				color = grey
			case city.PowerPlant:
				chr = "p"
				color = yellow
			case city.Dirt:
				chr = "░"
				color = yellow
			}
			c1, err := colorful.Hex(color)
			if err != nil {
				log.Fatal(err)
			}
			h, c, l := c1.Hcl()
			if cell == city.Dirt {
				l = 0.5 + (.125 - rand.Float64()/4)
			} else {
				l = r.jobs.Grid[y][x]
			}
			c2 := colorful.Hcl(h, c, l).Clamped()
			rowOutput = append(rowOutput, termenv.String(chr).Foreground(p.Color(c2.Hex())).String())
		}
		rows = append(rows, rowOutput)
	}
	return rows
}
