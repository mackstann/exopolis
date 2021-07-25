package city

type Row []Cell

type City []Row // TODO privatize some things

func NewCity(size int) *City {
	city := make(City, 0, size)
	for i := 0; i < size; i++ {
		city = append(city, make(Row, size))
	}
	return &city
}
