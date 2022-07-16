package component

type ScreenProvider struct {
	Screen ScreenFunType
}

func NewScreenApp() *ScreenProvider {
	return &ScreenProvider{Screen: InitScreen()}
}

func (s *ScreenProvider) Start(width int, height int, runtimeChan chan bool, snakes *snake, score *int, foodPoint *scope) error {
	return s.Screen(width, height, runtimeChan, snakes, score, foodPoint)
}
