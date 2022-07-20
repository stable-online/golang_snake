package component

//screenApp 屏幕应用
type screenApp struct {
	Screen screenFunType
}

//newScreenApp 屏幕实例化
func newScreenApp() *screenApp {
	return &screenApp{Screen: initScreenHandle()}
}

//start 启动屏幕渲染
func (s *screenApp) start(gameData *game) error {
	return s.Screen(gameData)
}
