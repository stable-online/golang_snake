package component

type ScreenProvider struct {
	Screen ScreenFunType
	snakes *snake
}

func NewScreenApp() *ScreenProvider {
	return &ScreenProvider{Screen: InitScreen(), snakes: new(snake)}
}

func (s *ScreenProvider) Start(width int, height int, runtimeChan chan bool, score *int, foodPoint *scope, direction int) error {
	return s.Screen(width, height, runtimeChan, s.snakes, score, foodPoint, direction)
}
