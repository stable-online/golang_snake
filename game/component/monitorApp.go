package component

type monitorApp struct {
	Monitor monitorFunType
}

//newMonitorApp 初始化监控应用
func newMonitorApp() *monitorApp {
	return &monitorApp{Monitor: initMonitor()}
}

//handle 启动监听
func (s *monitorApp) handle(game *game) {
	s.Monitor(game.getControl().getMoveChan(), game.getControl().getQuitChan())
}
