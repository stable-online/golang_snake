package game

import . "snake/game/component"

//gameService
type gameService struct {
	games GamesType
}

//init struct
func newGameService() *gameService {
	return &gameService{games: InitGames(NewScreenApp(), NewMonitorApp())}
}

//start
func (g *gameService) start() {
	g.games()
}
