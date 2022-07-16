package component

//main
const (
	RIGHT = 1 + iota
	LEFT
	UP
	DOWN
	QUIT
)

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

//gameData game 数据
type gameData struct {
	keyboardChan chan int
	quitChan     chan int
	runtimeChan  chan bool
	gameOver     bool
	score        int
	direction    int
}

//screenData 屏幕相关参数
type screenData struct {
	snakes    *snake
	foodPoint *scope
}
