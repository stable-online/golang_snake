package component

import (
	"errors"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"math/rand"
	"strconv"
	"time"
)

//screenFunType 启动屏幕展示
type screenFunType func(*game) error

//snakeFunType
type snakeFunType func(*game)

//foodFunType
type foodFunType func(int, int, *scope)

//moveFunType
type moveFunType func(int, int, chan bool, *snake, *int, *scope)

//initFood
func initFood() func(width int, height int, foodPoint *scope) {
	return func(width int, height int, foodPoint *scope) {
		if foodPoint.x == 0 && foodPoint.y == 0 {
			genFood(width, height, foodPoint)
		}
	}
}

//genFood
func genFood(width int, height int, foodPoint *scope) {
	foodPoint.x = generateRandInt(1, width-1)
	foodPoint.y = generateRandInt(4, height-1)
}

//handle
func handle(initSnake snakeFunType, initFood foodFunType, move moveFunType) screenFunType {
	return func(game *game) error {

		//游戏屏幕数据
		screen := game.getScreen()

		//动态设置表格(游戏中, 边框变动会进行动态调整)
		screen.initScreenSize()

		//初始化蛇
		initSnake(game)

		//初始化食物
		initFood(screen.getWidth(), screen.getHeight(), screen.getFoodPoint())

		//初始化移动
		move(screen.getWidth(), screen.getHeight(), game.getControl().getSnakeStatusChan(), screen.getSnakes(), screen.getScore(), screen.getFoodPoint())

		//拿到数据后,进行渲染界面
		return render(screen.getWidth(), screen.getHeight(), screen.getSnakes(), screen.getScore(), screen.getFoodPoint())
	}
}

//render
func render(width int, height int, snakes *snake, score *int, foodPoint *scope) error {

	if err := termbox.Clear(termbox.ColorDefault, termbox.ColorBlack); err != nil {
		return errors.New(err.Error())
	}

	//middle number
	var midWidth = width/2 - 8

	//setting the title
	for _, s := range "Snake games" {
		termbox.SetCell(midWidth, 1, s, termbox.ColorLightRed, termbox.ColorBlack)
		midWidth += runewidth.RuneWidth(s)
	}

	//setting the score
	x := 2
	for _, si := range "score:" {
		termbox.SetCell(x, 2, si, termbox.ColorLightRed, termbox.ColorBlack)
		x += runewidth.RuneWidth(si)
	}
	for _, sii := range strconv.Itoa(*score) {
		termbox.SetCell(x, 2, sii, termbox.ColorLightRed, termbox.ColorBlack)
		x += runewidth.RuneWidth(sii)
	}

	//setting quit tips
	wTip := width - 14 - 2
	for _, si := range "press esc quit" {
		termbox.SetCell(wTip, 2, si, termbox.ColorLightRed, termbox.ColorBlack)
		wTip += runewidth.RuneWidth(si)
	}

	//set frame
	w := 0
	for w <= width {
		termbox.SetCell(w, 3, ' ', termbox.ColorGreen, termbox.ColorLightGreen)
		termbox.SetCell(w, height, ' ', termbox.ColorGreen, termbox.ColorLightGreen)
		w += runewidth.RuneWidth(' ')
	}
	h := 0
	for h <= height {
		termbox.SetCell(0, h, ' ', termbox.ColorGreen, termbox.ColorLightGreen)
		termbox.SetCell(width, h, ' ', termbox.ColorGreen, termbox.ColorLightGreen)
		h += runewidth.RuneWidth(' ')
	}

	//setting snake
	for _, body := range snakes.snakeBody {
		termbox.SetCell(body.x, body.y, ' ', termbox.ColorLightRed, termbox.ColorLightRed)
	}

	//setting food
	termbox.SetCell(foodPoint.x, foodPoint.y, '@', termbox.ColorLightRed, termbox.ColorDefault)

	return termbox.Flush()
}

// initMove initMonitor for user keyboard
func initMove() moveFunType {
	return func(width int, height int, runtimeChan chan bool, snakes *snake, score *int, foodPoint *scope) {
		move(width, height, runtimeChan, snakes, score, foodPoint)
	}
}

//isDeath
func isDeath(width int, height int, snakes *snake) bool {

	s := head(snakes)

	return s.x >= width || s.y >= height || s.x <= 0 || s.y <= 3
}

//move
func move(width int, height int, runtimeChan chan bool, snakes *snake, score *int, foodPoint *scope) {
	scopes := head(snakes)

	switch snakes.direction {
	case RIGHT:
		scopes.x++
	case LEFT:
		scopes.x--
	case UP:
		scopes.y--
	case DOWN:
		scopes.y++
	}

	if isDeath(width, height, snakes) {
		runtimeChan <- true
		return
	}

	if scopes.x == foodPoint.x && scopes.y == foodPoint.y {
		snakes.len++
		*score++

		genFood(width, height, foodPoint)
	}

	if snakes.len > len(snakes.snakeBody) {
		snakes.snakeBody = append(snakes.snakeBody, scopes)
	} else {
		snakes.snakeBody = append(snakes.snakeBody[1:], scopes)
	}
}

//head
func head(snakes *snake) scope {
	return snakes.snakeBody[len(snakes.snakeBody)-1]
}

//generateRandInt
func generateRandInt(min, max int) int {
	rand.Seed(time.Now().Unix())

	n := max - min
	if n > 0 {
		return rand.Intn(n) + min
	}

	return 1
}

//initSnake snake
func initSnake() snakeFunType {
	return func(gameData *game) {

		//屏幕信息
		g := gameData.getScreen()

		//蛇的信息
		s := g.getSnakes()

		//如果蛇不存在, 需要初始化
		if len(s.getSnakeBody()) == 0 {

			//进行初始化body
			s.setSnakeBody(append(s.getSnakeBody(), scope{5, g.getHeight() - 2}))
			s.setSnakeBody(append(s.getSnakeBody(), scope{5, g.getHeight() - 3}))

			//初始化蛇的长度
			s.setLen(2)
		}

		//设置蛇跑的方向
		s.setDirection(gameData.getControl().getDirection())
	}
}

//initScreenHandle 初始化屏幕
func initScreenHandle() screenFunType {
	return handle(initSnake(), initFood(), initMove())
}
