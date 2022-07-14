package game

type game struct {
	games func()
}

func (g game) init() *game {
	return &game{games(screen(initSnake(), initFood(), initMove()), monitorKeyboard())}
}

func (g *game) start() {
	g.games()
}
