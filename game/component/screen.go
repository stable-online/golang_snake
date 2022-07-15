package component

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"math/rand"
	"strconv"
	"time"
)

//ScreenType 启动屏幕展示
type ScreenType func(int, int, chan bool)

//snakeType
type snakeType func(int, int)

//foodType
type foodType func(int, int)

//moveType
type moveType func(int, int, chan bool)

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

var (
	score     = 0
	foodPoint scope
	snakes    snake
)

//initFood
func initFood() func(width int, height int) {
	return func(width int, height int) {
		if foodPoint.x == 0 && foodPoint.y == 0 {
			genFood(width, height)
		}
	}
}

//genFood
func genFood(width int, height int) {
	foodPoint.x = generateRandInt(1, width-1)
	foodPoint.y = generateRandInt(4, height-1)
}

//screen
func screen(initSnake snakeType, initFood foodType, move moveType) ScreenType {
	return func(width int, height int, runtimeChan chan bool) {

		//init snakes
		initSnake(width, height)

		//init initFood
		initFood(width, height)

		//init move
		move(width, height, runtimeChan)

		render(width, height)
	}
}

//render
func render(width int, height int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorBlack)

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
	for _, sii := range strconv.Itoa(score) {
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
	termbox.Flush()
}

// initMove InitMonitor for user keyboard
func initMove() moveType {
	return func(width int, height int, runtimeChan chan bool) {
		move(width, height, runtimeChan)
	}
}

//isDeath
func isDeath(width int, height int) bool {

	s := head()

	return s.x >= width || s.y >= height || s.x <= 0 || s.y <= 3
}

//move
func move(width int, height int, runtimeChan chan bool) {
	scopes := head()

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

	if isDeath(width, height) {
		runtimeChan <- true
		return
	}

	if scopes.x == foodPoint.x && scopes.y == foodPoint.y {
		snakes.len++
		score++

		genFood(width, height)
	}

	if snakes.len > len(snakes.snakeBody) {
		snakes.snakeBody = append(snakes.snakeBody, scopes)
	} else {
		snakes.snakeBody = append(snakes.snakeBody[1:], scopes)
	}
}

//head
func head() scope {
	return snakes.snakeBody[len(snakes.snakeBody)-1]
}

//generateRandInt
func generateRandInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

//initSnake snake
func initSnake() snakeType {
	return func(width int, height int) {

		if len(snakes.snakeBody) == 0 {
			snakes.snakeBody = append(snakes.snakeBody, scope{5, height - 2})
			snakes.snakeBody = append(snakes.snakeBody, scope{5, height - 3})

			snakes.direction = UP
			snakes.len = 2
		}
	}
}

//InitScreen 初始化屏幕信息
func InitScreen() ScreenType {
	return screen(initSnake(), initFood(), initMove())
}
