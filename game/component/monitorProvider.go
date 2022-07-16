package component

type MonitorProvider struct {
	Monitor MonitorFunType
}

func NewMonitorApp() *MonitorProvider {
	return &MonitorProvider{Monitor: InitMonitor()}
}

func (s *MonitorProvider) Start(keyboardChan chan int, quitChan chan int) {
	s.Monitor(keyboardChan, quitChan)
}
