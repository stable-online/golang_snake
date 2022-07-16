package component

type ScreenApp struct {
	Screen ScreenFunType
}

func NewScreenApp() *ScreenApp {
	return &ScreenApp{Screen: InitScreen()}
}

func (s *ScreenApp) Start(width int, height int, runtimeChan chan bool, snakes *snake, score *int, foodPoint *scope) error {
	return s.Screen(width, height, runtimeChan, snakes, score, foodPoint)
}
