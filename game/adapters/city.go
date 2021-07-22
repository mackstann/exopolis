package adapters

import (
	cityDomain "github.com/mackstann/exopolis/city/domain"
)

type CityAdapter struct {
	city *cityDomain.City
}

func NewCityAdapter(city *cityDomain.City) *CityAdapter {
	return &CityAdapter{
		city: city,
	}
}

func (a *CityAdapter) Get() *cityDomain.City {
	return a.city
}
