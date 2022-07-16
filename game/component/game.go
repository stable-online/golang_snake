package component

import (
	"github.com/nsf/termbox-go"
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

// GamesType defined method
type GamesType func(screen ScreenFunType, monitorKeyboard MonitorFunType)

// InitGames 游戏初始化高阶函数
func InitGames() GamesType {
	return func(screen ScreenFunType, monitor MonitorFunType) {

		//keyboardChan channel
		var (
			keyboardChan = make(chan int)
			quitChan     = make(chan int)
			runtimeChan  = make(chan bool, 1)
			gameOver     = false
			snakes       snake
		)

		//init box
		if initErr := termbox.Init(); initErr != nil {
			panic(initErr)
		}

		//box to close
		defer termbox.Close()

		//monitor keyboardChan
		go monitor(keyboardChan, quitChan)

		for {
			select {
			case operator := <-keyboardChan:
				if operator != 0 {
					snakes.direction = operator
				}
			case <-quitChan:
				return
			case msg := <-runtimeChan:
				gameOver = msg
			default:
				if !gameOver {
					width, height := termbox.Size()
					if err := screen(width-1, height-1, runtimeChan, &snakes); err != nil {
						panic(err.Error())
					}
				}
				flush()
			}
		}
	}
}

//flush
func flush() {
	time.Sleep(time.Duration(100-(score/10)) * time.Millisecond)
}
