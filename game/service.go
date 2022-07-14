package game

import "snake/game/component"

//service
type service struct {
	games   component.GamesType      //游戏主控
	screen  component.ScreenType     //屏幕刷新
	monitor component.MonitorKeyType //键盘控制
}

//start
func (g *service) start() {
	g.games(g.screen, g.monitor)
}

//init struct
func newGame() *service {
	return &service{games: component.InitGames(), screen: component.InitScreen(), monitor: component.InitMonitor()}
}
