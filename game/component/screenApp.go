package component

//screenApp 屏幕应用
type screenApp struct {
	screen screenFunType
}

//newScreenApp 初始化屏幕应用
func newScreenApp() *screenApp {
	return &screenApp{screen: initScreenHandle()}
}

//start 启动屏幕渲染
func (s *screenApp) start(gameData *game) error {
	return s.screen(gameData)
}
