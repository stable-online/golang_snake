package component

import (
	"github.com/nsf/termbox-go"
)

//gameService  game 服务
type gameService struct {
	screenApp  *screenProvider
	monitorApp *monitorProvider
}

//NewGameService 实例化游戏服务
func NewGameService() *gameService {
	return &gameService{screenApp: newScreenApp(), monitorApp: newMonitorApp()}
}

//Start 开始游戏
func (g *gameService) Start(data *gameData) {

	if initErr := termbox.Init(); initErr != nil {
		panic(initErr)
	}

	//box to close
	defer termbox.Close()

	//monitor keyboardChan
	go g.monitorApp.start(data)

	//run job
	g.run(data)
}

//run 进行
func (g *gameService) run(data *gameData) {
	for {
		select {
		case operator := <-data.keyboardChan:
			data.direction = operator
		case <-data.quitChan:
			return
		case msg := <-data.runtimeChan:
			data.gameOver = msg
		default:
			if !data.gameOver {
				width, height := termbox.Size()
				if err := g.screenApp.start(width-1, height-1, data); err != nil {
					panic(err.Error())
				}
			}
			flush(data.score)
		}
	}
}
