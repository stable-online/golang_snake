package component

type monitorProvider struct {
	Monitor monitorFunType
}

func NewMonitorApp() *monitorProvider {
	return &monitorProvider{Monitor: initMonitor()}
}

func (s *monitorProvider) Start(GameData *gameData) {
	s.Monitor(GameData.keyboardChan, GameData.quitChan)
}
