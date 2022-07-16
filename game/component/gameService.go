package component

import (
	"github.com/nsf/termbox-go"
)

//GameService  game 服务
type GameService struct {
	screenApp  *ScreenProvider
	monitorApp *monitorProvider
	data       *gameData
}

func NewGameService() *GameService {
	return &GameService{screenApp: newScreenApp(), monitorApp: NewMonitorApp(), data: newGameData()}
}

func (g *GameService) Start() {

	if initErr := termbox.Init(); initErr != nil {
		panic(initErr)
	}

	//box to close
	defer termbox.Close()

	//monitor keyboardChan
	go g.monitorApp.Start(g.data)

	//run job
	g.working()
}

func (g *GameService) working() {
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
				if err := g.screenApp.start(width-1, height-1, g.data); err != nil {
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
