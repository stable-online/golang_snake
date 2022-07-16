package component

type MonitorApp struct {
	Monitor MonitorFunType
}

func NewMonitorApp() *MonitorApp {
	return &MonitorApp{Monitor: InitMonitor()}
}

func (s *MonitorApp) Start(keyboardChan chan int, quitChan chan int) {
	s.Monitor(keyboardChan, quitChan)
}
