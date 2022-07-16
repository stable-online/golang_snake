package component

type MonitorProvider struct {
	Monitor monitorFunType
}

func NewMonitorApp() *MonitorProvider {
	return &MonitorProvider{Monitor: initMonitor()}
}

func (s *MonitorProvider) Start(GameData *gameData) {
	s.Monitor(GameData.keyboardChan, GameData.quitChan)
}
