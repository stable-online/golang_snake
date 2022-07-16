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

//GameData game 数据
type GameData struct {
	keyboardChan chan int
	quitChan     chan int
	runtimeChan  chan bool
	gameOver     bool
	score        int
	direction    int
}

//GameService  game 服务
type GameService struct {
	screenApp  *ScreenProvider
	monitorApp *MonitorProvider
	data       *GameData
	score      int
}
