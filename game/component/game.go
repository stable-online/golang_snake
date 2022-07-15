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
					screen(width-1, height-1, runtimeChan)
				}
				flushScreen()
			}
		}
	}
}

//flushScreen
func flushScreen() {
	time.Sleep(time.Duration(100-(score/10)) * time.Millisecond)
}
