package component

type screenProvider struct {
	Screen     screenFunType
	screenData *screenData
}

func newScreenApp() *screenProvider {
	return &screenProvider{Screen: initScreen(), screenData: newScreenData()}
}

func (s *screenProvider) start(width int, height int, gameData *gameData) error {
	return s.Screen(width, height, gameData.runtimeChan, s.screenData.snakes, &gameData.score, s.screenData.foodPoint, gameData.direction)
}
