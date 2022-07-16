package component

type MonitorProvider struct {
	Monitor MonitorFunType
}

func NewMonitorApp() *MonitorProvider {
	return &MonitorProvider{Monitor: InitMonitor()}
}

func (s *MonitorProvider) Start(quitChan chan int, GameData *GameData) {
	s.Monitor(GameData.keyboardChan, quitChan)
}
