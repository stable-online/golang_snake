package component

import (
	"github.com/nsf/termbox-go"
)

func NewGameService() *GameService {
	return &GameService{screenApp: NewScreenApp(), monitorApp: NewMonitorApp(), data: NewGameData(), score: 0}
}

func (g *GameService) Start() {

	//keyboardChan channel
	var (
		foodPoint scope
		direction = UP
	)

	//init box
	if initErr := termbox.Init(); initErr != nil {
		panic(initErr)
	}

	//box to close
	defer termbox.Close()

	//monitor keyboardChan
	go g.monitorApp.Start(g.data)

	for {
		select {
		case operator := <-g.data.keyboardChan:
			if operator != 0 {
				direction = operator
			}
		case <-g.data.quitChan:
			return
		case msg := <-g.data.runtimeChan:
			g.data.gameOver = msg
		default:
			if !g.data.gameOver {
				width, height := termbox.Size()
				if err := g.screenApp.Start(width-1, height-1, g.data.runtimeChan, &g.data.score, &foodPoint, direction); err != nil {
					panic(err.Error())
				}
			}
			g.flush(&g.data.score)
		}
	}
}

func (g *GameService) flush(score *int) {
	Flush(*score)
}
