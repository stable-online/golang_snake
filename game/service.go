package game

import "snake/game/component"

//service
type service struct {
	games   component.GamesType      //游戏主控
	screen  component.ScreenFunType  //屏幕刷新
	monitor component.MonitorFunType //键盘控制
}

//init struct
func newGame() *service {
	return &service{games: component.InitGames(), screen: component.InitScreen(), monitor: component.InitMonitor()}
}

//start
func (g *service) start() {
	g.games(g.screen, g.monitor)
}
