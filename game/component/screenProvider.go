package component

type ScreenProvider struct {
	Screen     screenFunType
	screenData *screenData
}

func newScreenApp() *ScreenProvider {
	return &ScreenProvider{Screen: initScreen(), screenData: newScreenData()}
}

func (s *ScreenProvider) start(width int, height int, gameData *gameData) error {
	return s.Screen(width, height, gameData.runtimeChan, s.screenData.snakes, &gameData.score, s.screenData.foodPoint, gameData.direction)
}
