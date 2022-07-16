package component

import (
	"github.com/nsf/termbox-go"
	"time"
)

// GamesType defined method
type GamesType func()

// InitGames 游戏初始化
func InitGames(screenApp *ScreenApp, monitorApp *MonitorApp) GamesType {
	return func() {

		//keyboardChan channel
		var (
			keyboardChan = make(chan int)
			quitChan     = make(chan int)
			runtimeChan  = make(chan bool, 1)
			gameOver     = false
			snakes       snake
			score        = 0
			foodPoint    scope
		)

		//init box
		if initErr := termbox.Init(); initErr != nil {
			panic(initErr)
		}

		//box to close
		defer termbox.Close()

		//monitor keyboardChan
		go monitorApp.Start(keyboardChan, quitChan)

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
					if err := screenApp.Start(width-1, height-1, runtimeChan, &snakes, &score, &foodPoint); err != nil {
						panic(err.Error())
					}
				}
				flush(&score)
			}
		}
	}
}

//flush
func flush(score *int) {
	time.Sleep(time.Duration(100-(*score/10)) * time.Millisecond)
}
