package component

type ScreenProvider struct {
	Screen    ScreenFunType
	snakes    *snake
	foodPoint *scope
}

func NewScreenApp() *ScreenProvider {
	return &ScreenProvider{Screen: initScreen(), snakes: new(snake), foodPoint: new(scope)}
}

func (s *ScreenProvider) Start(width int, height int, data *GameData) error {
	return s.Screen(width, height, data.runtimeChan, s.snakes, &data.score, s.foodPoint, data.direction)
}
