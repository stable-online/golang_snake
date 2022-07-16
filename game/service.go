package game

import "snake/game/component"

//service
type service struct {
	games component.GamesType
}

//init struct
func newGame() *service {
	return &service{games: component.InitGames(component.InitScreen(), component.InitMonitor())}
}

//start
func (g *service) start() {
	g.games()
}
