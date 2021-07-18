package renderer

import (
	"fmt"

	"github.com/mackstann/exoplanet/city/domain"

	"github.com/muesli/termenv"
)

func Render(city domain.City, n *domain.JobTransportNetwork) {
	for _, row := range textualize(city, n) {
		fmt.Println(row)
	}
}

func textualize(city domain.City, n *domain.JobTransportNetwork) []string {
	rows := make([]string, 0, len(city))
	for y, row := range city {
		rowOutput := ""
		for x, cell := range row {
			nCell := n.Grid[y][x]
			//fmt.Printf("%f ", nCell.Temperature)
			intensity := ""
			if nCell.Temperature < 0.2 {
				intensity = "55"
			} else if nCell.Temperature < 0.4 {
				intensity = "77"
			} else if nCell.Temperature < 0.6 {
				intensity = "99"
			} else if nCell.Temperature < 0.8 {
				intensity = "bb"
			} else {
				intensity = "dd"
			}
			c := "."
			color := ""
			switch cell.Typ {
			case domain.House:
				c = "h"
				color = intensity + "0000"
				//fmt.Printf("house %s\n", color)
			case domain.Farm:
				c = "f"
				color = "00" + intensity + "00"
				//fmt.Printf("farm %s\n", color)
			case domain.Road:
				c = "r"
				color = intensity + intensity + intensity
				//fmt.Printf("road %s\n", color)
			}
			p := termenv.ColorProfile()
			if intensity != "" && c != " " && len(color) == 6 {
				rowOutput += termenv.String(c).Foreground(p.Color("#" + color)).String()
			} else {
				rowOutput += c
			}
		}
		//fmt.Printf("\n")
		rows = append(rows, rowOutput)
	}
	return rows
}
