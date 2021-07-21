package game

type City interface{}

type CityService interface {
	get() *City
}

type MapGeneratorService interface {
	generate(*City)
}

type GameLoopService interface {
	run(*City)
}

type Game struct {
	city         CityService
	mapGenerator MapGeneratorService
	gameLoop     GameLoopService
}

func NewGame(city CityService, mapGenerator MapGeneratorService, gameLoop GameLoopService) *Game {
	return &Game{
		city:         city,
		mapGenerator: mapGenerator,
		gameLoop:     gameLoop,
	}
}

func (g *Game) run() {
	city := g.city.get()
	g.mapGenerator.generate(city)
	g.gameLoop.run(city)
}
