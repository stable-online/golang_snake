package game

import "snake/game/component"

//gameService
type gameService struct {
	games component.GamesType
}

//init struct
func newGameService() *gameService {
	return &gameService{games: component.InitGames(NewScreenService().start(), NewMonitorService().start())}
}

//start
func (g *gameService) start() {
	g.games()
}
