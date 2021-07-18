package renderer

import (
	"fmt"
	"log"

	"github.com/mackstann/exoplanet/city/domain"

	"github.com/muesli/termenv"
)

func Render(city domain.City, n *domain.JobTransportNetwork) {
	if termenv.ColorProfile() != termenv.TrueColor {
		log.Fatalf("not enough color! %v, want %v", termenv.ColorProfile(), termenv.TrueColor)
	}

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
			fmt.Printf("TEMP %f ", nCell.Temperature)
			temp255 := int(nCell.Temperature * 255)
			intensity := fmt.Sprintf("%02x", temp255)
			fmt.Println(nCell.Temperature)
			fmt.Printf("intensity %s\n", intensity)
			c := "."
			color := ""
			switch cell.Typ {
			case domain.House:
				c = "h"
				color = intensity + "0000"
				fmt.Printf("house %s\n", color)
			case domain.Farm:
				c = "f"
				color = "00" + intensity + "00"
				fmt.Printf("farm %s\n", color)
			case domain.Road:
				c = "r"
				color = intensity + intensity + intensity
				fmt.Printf("road %s\n", color)
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
