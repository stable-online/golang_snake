package component

import (
	"github.com/nsf/termbox-go"
)

//GameService  game 服务
type GameService struct {
	screenApp  *ScreenProvider
	monitorApp *MonitorProvider
	data       *GameData
}

func NewGameService() *GameService {
	return &GameService{screenApp: NewScreenApp(), monitorApp: NewMonitorApp(), data: NewGameData()}
}

func (g *GameService) Start() {

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
				g.data.direction = operator
			}
		case <-g.data.quitChan:
			return
		case msg := <-g.data.runtimeChan:
			g.data.gameOver = msg
		default:
			if !g.data.gameOver {
				width, height := termbox.Size()
				if err := g.screenApp.Start(width-1, height-1, g.data); err != nil {
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
