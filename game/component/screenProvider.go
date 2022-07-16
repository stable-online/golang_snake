package component

type ScreenProvider struct {
	Screen    screenFunType
	snakes    *snake
	foodPoint *scope
}

func newScreenApp() *ScreenProvider {
	return &ScreenProvider{Screen: initScreen(), snakes: new(snake), foodPoint: new(scope)}
}

func (s *ScreenProvider) start(width int, height int, data *gameData) error {
	return s.Screen(width, height, data.runtimeChan, s.snakes, &data.score, s.foodPoint, data.direction)
}
