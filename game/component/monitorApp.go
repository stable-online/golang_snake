package component

type monitorApp struct {
	Monitor monitorFunType
}

//newMonitorApp 实例化
func newMonitorApp() *monitorApp {
	return &monitorApp{Monitor: initMonitor()}
}

//handle 启动监听
func (s *monitorApp) handle(game *game) {
	s.Monitor(game.getControl().getMoveChan(), game.getControl().getQuitChan())
}
