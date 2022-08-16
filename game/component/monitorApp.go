package component

type monitorApp struct {
	monitor monitorFunType
}

//newMonitorApp 初始化监控应用
func newMonitorApp() *monitorApp {
	return &monitorApp{monitor: initMonitor()}
}

//handle 启动监听
func (s *monitorApp) handle(game *game) {
	s.monitor(game.getControl().getMoveChan(), game.getControl().getQuitChan())
}
