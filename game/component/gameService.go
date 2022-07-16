package component

type GameService struct {
	games GamesType
}

func NewGameService() *GameService {
	return &GameService{games: InitGames(NewScreenApp(), NewMonitorApp())}
}

func (g *GameService) Start() {
	g.games(g)
}
