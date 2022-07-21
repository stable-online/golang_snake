package component

import "github.com/nsf/termbox-go"

//scope
type scope struct {
	x int
	y int
}

//snake
type snake struct {
	snakeBody []scope
	direction int
	len       int
}

//setLen 蛇的长度
func (s *snake) setLen(len int) {
	s.len = len
}

//setSnakeBody 设置蛇的身体
func (s *snake) setSnakeBody(snakeBody []scope) {
	s.snakeBody = snakeBody
}

//setDirection 设置蛇要走的方向
func (s *snake) setDirection(direction int) {
	s.direction = direction
}

//getSnakeBody 获取蛇的身体
func (s snake) getSnakeBody() []scope {
	return s.snakeBody
}

//getDirection 方向
func (s snake) getDirection() int {
	return s.direction
}

//getLen 长度
func (s snake) getLen() int {
	return s.len
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//screen 屏幕相关参数
type screen struct {
	snakes    *snake
	foodPoint *scope
	width     int
	height    int
	score     int
}

//newScreen 实例化屏幕
func newScreen() *screen {
	return &screen{snakes: new(snake), foodPoint: new(scope), score: 0}
}

//getScore 获取分数
func (s *screen) getScore() *int {
	return &s.score
}

//setScore  设置得分
func (s *screen) setScore(score int) {
	s.score = score
}

//setWidth 设置边框的宽度
func (s *screen) setWidth(width int) {
	s.width = width
}

//setHeight 设置边框的高度
func (s *screen) setHeight(height int) {
	s.height = height
}

//initScreenSize 实时获取盒子尺寸
func (s *screen) initScreenSize() {

	//获取盒子尺寸
	size, height := termbox.Size()

	//设置边框的宽
	s.setWidth(size - 1)

	//设置边框的高
	s.setHeight(height - 1)
}

//getSnakes 获取蛇
func (s screen) getSnakes() *snake {
	return s.snakes
}

//getFoodPoint 获取食物的坐标
func (s screen) getFoodPoint() *scope {
	return s.foodPoint
}

//getWidth 获取屏幕的狂赌
func (s screen) getWidth() int {
	return s.width
}

//getHeight 获取屏幕的高度
func (s screen) getHeight() int {
	return s.height
}

//运行时数据
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//control 控制
type control struct {
	moveChannel        chan int
	quitChannel        chan int
	snakeStatusChannel chan bool
	gameOver           bool
	direction          int
}

//newControl 实例化控制器
func newControl() *control {
	return &control{moveChannel: make(chan int), quitChannel: make(chan int), snakeStatusChannel: make(chan bool, 1), gameOver: false, direction: UP}
}

//setGameOver 设置蛇已死亡
func (r *control) setGameOver(gameOver bool) {
	r.gameOver = gameOver
}

//setDirection 设置玩家方向指令
func (r *control) setDirection(direction int) {
	r.direction = direction
}

//getMoveChan 键盘监听--蛇移动管道
func (r *control) getMoveChan() chan int {
	return r.moveChannel
}

//getQuitChan 键盘监听--退出管道
func (r *control) getQuitChan() chan int {
	return r.quitChannel
}

//getSnakeStatusChan 游戏状态管道
func (r *control) getSnakeStatusChan() chan bool {
	return r.snakeStatusChannel
}

//getActivity 如果蛇还在活跃中
func (r *control) getActivity() bool {
	return r.gameOver == false
}

//getDirection 获取玩家的方向的指令
func (r *control) getDirection() int {
	return r.direction
}

//游戏数据总览数据
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//game control 游戏数据结构
type game struct {
	//控制
	control *control
	//屏幕
	screen *screen
}

// getControl 控制信息
func (g *game) getControl() *control {
	return g.control
}

//getScreen 屏幕信息
func (g *game) getScreen() *screen {
	return g.screen
}

//NewGameData 实例化
func NewGameData() *game {
	return &game{control: newControl(), screen: newScreen()}
}
