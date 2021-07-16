package renderer

import (
	"fmt"

	"github.com/mackstann/exoplanet/city/domain"
)

func Render(city domain.City) {
	for _, row := range textualize(city) {
		fmt.Println(row)
	}
}

func textualize(city domain.City) []string {
	rows := make([]string, 0, len(city))
	for _, row := range city {
		rowOutput := make([]rune, 0, len(row))
		for _, cell := range row {
			var c rune = ' '
			switch cell.Typ {
			case domain.House:
				c = 'h'
			case domain.Farm:
				c = 'f'
			}
			rowOutput = append(rowOutput, c)
		}
		rows = append(rows, string(rowOutput))
	}
	return rows
}
