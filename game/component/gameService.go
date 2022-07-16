package component

import (
	"github.com/nsf/termbox-go"
)

func NewGameService() *GameService {
	return &GameService{screenApp: NewScreenApp(), monitorApp: NewMonitorApp(), data: NewGameData()}
}

func (g *GameService) Start() {

	//keyboardChan channel
	var (
		quitChan    = make(chan int)
		runtimeChan = make(chan bool, 1)
		gameOver    = false
		snakes      snake
		score       = 0
		foodPoint   scope
	)

	//init box
	if initErr := termbox.Init(); initErr != nil {
		panic(initErr)
	}

	//box to close
	defer termbox.Close()

	//monitor keyboardChan
	go g.monitorApp.Start(quitChan, g.data)

	for {
		select {
		case operator := <-g.data.keyboardChan:
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
				if err := g.screenApp.Start(width-1, height-1, runtimeChan, &snakes, &score, &foodPoint); err != nil {
					panic(err.Error())
				}
			}
			g.flush(&score)
		}
	}
}

func (g *GameService) flush(score *int) {
	Flush(*score)
}
