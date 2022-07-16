package component

type ScreenProvider struct {
	Screen    ScreenFunType
	snakes    *snake
	foodPoint *scope
}

func NewScreenApp() *ScreenProvider {
	return &ScreenProvider{Screen: InitScreen(), snakes: new(snake), foodPoint: new(scope)}
}

func (s *ScreenProvider) Start(width int, height int, runtimeChan chan bool, score *int, direction int) error {
	return s.Screen(width, height, runtimeChan, s.snakes, score, s.foodPoint, direction)
}
