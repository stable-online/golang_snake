package component

type MonitorProvider struct {
	Monitor MonitorFunType
}

func NewMonitorApp() *MonitorProvider {
	return &MonitorProvider{Monitor: InitMonitor()}
}

func (s *MonitorProvider) Start(GameData *GameData) {
	s.Monitor(GameData.keyboardChan, GameData.quitChan)
}
