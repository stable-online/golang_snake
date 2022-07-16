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

type GameService struct {
	screenApp  *ScreenProvider
	monitorApp *MonitorProvider
}
