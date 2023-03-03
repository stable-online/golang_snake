package component

import "time"

type Screen interface {
	// start
	//
	// @Description: 启动屏幕
	// @return error
	start() error

	// setDirection
	//
	// @Description:设置蛇的方向
	// @param operator
	setDirection(operator int)

	// getActivity
	//
	// @Description: 获取蛇是否还活着
	// @return bool
	getActivity() bool

	// getSnakeStatusChan
	//
	// @Description: 获取当前蛇的状态
	// @return <-chan
	getSnakeStatusChan() <-chan bool

	// setGameOver
	//
	// @Description: 设置游戏的状态
	// @param status
	setGameOver(status bool)

	// flush
	//
	// @Description: 刷新屏幕
	flush()
}

//screenApp 屏幕应用
type screenApp struct {
	//
	//  screen
	//  @Description: 屏幕
	//
	screen screenFunType
	//
	//  role
	//  @Description: 角色
	//
	role *game
}

//newScreenApp 初始化屏幕应用
func newScreenApp() Screen {
	return &screenApp{screen: initScreenHandle(), role: newGameData()}
}

//start 启动屏幕渲染
func (s *screenApp) start() error {
	return s.screen(s.role)
}

// setDirection
//
// @Description:
// @receiver s
// @param operator
func (s *screenApp) setDirection(operator int) {
	//给游戏控制设置方向
	s.role.getControl().setDirection(operator)
}

// getSnakeStatusChan
//
// @Description:
// @receiver s
// @return <-chan
func (s *screenApp) getSnakeStatusChan() <-chan bool {
	return s.role.getControl().getSnakeStatusChan()
}

// setGameOver
//
// @Description:
// @receiver s
// @param status
func (s *screenApp) setGameOver(status bool) {
	s.role.getControl().setGameOver(status)
}

// getActivity
//
// @Description:
// @receiver s
// @return bool
func (s *screenApp) getActivity() bool {
	return s.role.getControl().getActivity()
}

// flush
//
// @Description:
// @receiver s
func (s *screenApp) flush() {
	time.Sleep(time.Duration(150-(*s.role.getScreen().getScore()/10)) * time.Millisecond)
}
