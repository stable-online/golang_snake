package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"math/rand"
	"strconv"
	"time"
)

//main
const (
	RIGHT = 1 + iota
	LEFT
	UP
	DOWN
	QUIT
)

type scope struct {
	x int
	y int
}

type snake struct {
	snakeBody []scope
	direction int
	len       int
}

var (
	snakes    snake
	score     = 0
	foodPoint scope
	keyboard  = make(chan int)
	over      = false
	quit      = make(chan int)
)

func main() {

	//to create the game starter
	GameStarter := Games(Screen(initSnake(), initFood(), initMove()), MonitorKeyboard())

	//starting
	GameStarter()
}

// Games play
func Games(screen func(width int, height int), monitorKeyboard func()) func() {
	return func() {

		//init box
		if initErr := termbox.Init(); initErr != nil {
			panic(initErr)
		}

		//box to close
		defer termbox.Close()

		//monitor keyboard
		go monitorKeyboard()

		for {
			select {

			case operator := <-keyboard:
				switch operator {
				case operator:
					if operator != 0 {
						snakes.direction = operator
					}
				}
			case <-quit:
				return

			default:

				width, height := termbox.Size()

				screen(width-1, height-1)

				flush()
			}
		}
	}
}

func flush() {
	time.Sleep(time.Duration(100-(score/10)) * time.Millisecond)
}

func initFood() func(width int, height int) {
	return func(width int, height int) {
		if foodPoint.x == 0 && foodPoint.y == 0 {
			genFood(width, height)
		}
	}
}

func genFood(width int, height int) {
	foodPoint.x = GenerateRandInt(1, width-1)
	foodPoint.y = GenerateRandInt(4, height-1)
}

//Screen 区域
func Screen(initSnake func(snakes snake, width int, height int) snake, initFood func(width int, height int), move func(width int, height int)) func(width int, height int) {
	return func(width int, height int) {

		snakes = initSnake(snakes, width, height)

		//init initFood
		initFood(width, height)

		if !over {
			//init move
			move(width, height)
		}

		render(width, height)
	}
}

func render(width int, height int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	//middle number
	var midWidth = width/2 - 8

	//setting the title
	for _, s := range "Snake Games" {
		termbox.SetCell(midWidth, 1, s, termbox.ColorDefault, termbox.ColorDefault)
		midWidth += runewidth.RuneWidth(s)
	}

	//setting the score
	x := 2
	for _, si := range "score:" {
		termbox.SetCell(x, 2, si, termbox.ColorRed, termbox.ColorDefault)
		x += runewidth.RuneWidth(si)
	}
	for _, sii := range strconv.Itoa(score) {
		termbox.SetCell(x, 2, sii, termbox.ColorRed, termbox.ColorDefault)
		x += runewidth.RuneWidth(sii)
	}

	//setting quit tips
	wTip := width - 14 - 2
	for _, si := range "press esc quit" {
		termbox.SetCell(wTip, 2, si, termbox.ColorRed, termbox.ColorDefault)
		wTip += runewidth.RuneWidth(si)
	}

	//set frame
	w := 0
	for w <= width {
		termbox.SetCell(w, 3, '*', termbox.ColorGreen, termbox.ColorDefault)
		termbox.SetCell(w, height, '*', termbox.ColorGreen, termbox.ColorDefault)
		w += runewidth.RuneWidth('*')
	}
	h := 0
	for h <= height {
		termbox.SetCell(0, h, '*', termbox.ColorGreen, termbox.ColorDefault)
		termbox.SetCell(width, h, '*', termbox.ColorGreen, termbox.ColorDefault)
		h += runewidth.RuneWidth('*')
	}

	//setting snake
	for _, body := range snakes.snakeBody {
		termbox.SetCell(body.x, body.y, 'o', termbox.ColorLightGreen, termbox.ColorDefault)
	}

	//setting food
	termbox.SetCell(foodPoint.x, foodPoint.y, '0', termbox.ColorLightRed, termbox.ColorDefault)
	termbox.Flush()
}

//MonitorKeyboard  用户键盘控制
func MonitorKeyboard() func() {
	return func() {

		termbox.SetInputMode(termbox.InputEsc)

		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyArrowLeft:
					keyboard <- LEFT
				case termbox.KeyArrowDown:
					keyboard <- DOWN
				case termbox.KeyArrowRight:
					keyboard <- RIGHT
				case termbox.KeyArrowUp:
					keyboard <- UP
				case termbox.KeyEsc:
					quit <- QUIT
				}
			default:
				keyboard <- 0
			case termbox.EventError:
				panic(ev.Err)
			}
		}
	}
}

// initMove MonitorKeyboard for user keyboard
func initMove() func(width int, height int) {
	return func(width int, height int) {
		move(width, height)
	}
}

func isDeath(width int, height int) bool {

	s := head()

	return s.x >= width || s.y >= height || s.x <= 0 || s.y <= 3
}

func move(width int, height int) {
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
		over = true
		return
	}

	if scopes.x == foodPoint.x && scopes.y == foodPoint.y {
		snakes.len++
		score++

		foodPoint.x = GenerateRandInt(1, width-1)
		foodPoint.y = GenerateRandInt(4, height-1)
	}

	if snakes.len > len(snakes.snakeBody) {
		snakes.snakeBody = append(snakes.snakeBody, scopes)
	} else {
		snakes.snakeBody = append(snakes.snakeBody[1:], scopes)
	}
}

func head() scope {
	return snakes.snakeBody[len(snakes.snakeBody)-1]
}

func GenerateRandInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

//initSnake snake
func initSnake() func(snakes snake, width int, height int) snake {
	return func(snakes snake, width int, height int) snake {

		if len(snakes.snakeBody) == 0 {
			snakes.snakeBody = append(snakes.snakeBody, scope{5, height - 2})
			snakes.snakeBody = append(snakes.snakeBody, scope{5, height - 3})

			snakes.direction = UP
			snakes.len = 2
		}

		return snakes
	}
}
