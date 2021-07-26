package renderer

import (
	"fmt"
	"log"

	"github.com/mackstann/exopolis/city"

	"github.com/muesli/termenv"
)

func Render(city city.City, n *city.JobTransportNetwork) []string {
	// TODO belongs at game scope
	/*
		argggggggggggggggg
		if termenv.ColorProfile() != termenv.TrueColor {
			log.Fatalf("not enough color! %v, want %v", termenv.ColorProfile(), termenv.TrueColor)
		}
	*/
	log.Println("render it")
	return textualize(city, n)
}

func textualize(c city.City, n *city.JobTransportNetwork) []string {
	rows := make([]string, 0, len(c))
	for y, row := range c {
		rowOutput := ""
		for x, cell := range row {
			temp255 := int(c[y][x].Resources.Jobs * 255)
			intensity := fmt.Sprintf("%02x", temp255)
			c := "."
			color := ""
			switch cell.Typ {
			case city.House:
				c = "■"
				color = intensity + "0000"
			case city.Farm:
				c = "▤"
				color = "00" + intensity + "00"
			case city.Road:
				c = "▪"
				color = intensity + intensity + intensity
			case city.PowerPlant:
				c = "p"
				color = intensity + intensity + "00"
			case city.Dirt:
				c = "░"
				color = "0000" + intensity
			}
			p := termenv.ColorProfile()
			if intensity != "" && c != " " && len(color) == 6 {
				rowOutput += termenv.String(c).Foreground(p.Color("#" + color)).String()
			} else {
				rowOutput += c
			}
		}
		rows = append(rows, rowOutput)
	}
	return rows
}
