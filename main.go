package main

import (
	"fmt"
)

type CellType int

const (
	// Dirt, being the zero CellType value, is the default value of a newly allocated cell.
	Dirt CellType = 0

	Road      = 100
	PowerLine = 101

	House = 200
	Farm  = 201

	PowerPlant = 300
)

type Cell struct {
	typ CellType

	farm       *FarmCell
	house      *HouseCell
	powerPlant *PowerPlantCell
	powerLine  *PowerLineCell
	road       *RoadCell
}

type FarmCell struct {
	productivity int
	accessTo     struct {
		water   int
		workers int
	}
}

type HouseCell struct {
	population int
	accessTo   struct {
		electricity int
		food        int
		work        int
	}
}

type PowerLineCell struct{}

type PowerPlantCell struct {
	production int
}

type RoadCell struct{}

type Row []Cell

type City []Row

func NewCity(size int) City {
	city := make(City, 0, size)
	for i := 0; i < size; i++ {
		city = append(city, make(Row, size))
	}
	return city
}

func main() {
	city := NewCity(20)
	city[5][5].typ = House
	city[5][5].house = &HouseCell{
		population: 10,
	}
	for _, row := range render(city) {
		fmt.Println(row)
	}
}

func render(city City) []string {
	rows := make([]string, 0, len(city))
	for _, row := range city {
		rowOutput := make([]rune, 0, len(row))
		for _, cell := range row {
			var c rune = ' '
			if cell.typ != Dirt {
				c = 'x'
			}
			rowOutput = append(rowOutput, c)
		}
		rows = append(rows, string(rowOutput))
	}
	return rows
}
