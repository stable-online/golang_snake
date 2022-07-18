package component

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

//screenData 屏幕相关参数
type screenData struct {
	snakes    *snake
	foodPoint *scope
}

func newScreenData() *screenData {
	return &screenData{snakes: new(snake), foodPoint: new(scope)}
}
