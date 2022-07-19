package component

import (
	"github.com/nsf/termbox-go"
)

//gameService 定义游戏行为
type gameService struct {
	screenApp  *screenApp
	monitorApp *monitorApp
}

//NewGameService 实例化游戏服务
func NewGameService() *gameService {
	return &gameService{screenApp: newScreenApp(), monitorApp: newMonitorApp()}
}

//Start 开始游戏
func (g *gameService) Start(game *game) {

	//启动包
	if initErr := termbox.Init(); initErr != nil {
		panic(initErr)
	}

	//函数退出时, 关闭包
	defer termbox.Close()

	//加入协程空间, 异步监听用户点击事件
	go g.monitorApp.handle(game)

	//游戏启动
	g.run(game)
}

//run 游戏启动
func (g *gameService) run(game *game) {

	for {
		select {

		//键盘移动事件
		case operator := <-game.getControl().getMoveChan():
			//给游戏控制设置方向
			game.getControl().setDirection(operator)

		//点击ECS事件
		case <-game.getControl().getQuitChan():
			//退出刷新
			return

		//游戏角色状态
		case status := <-game.getControl().getPlayGameStatusChan():
			//设置人物状态
			game.getControl().setGameOver(status)

		default:

			//如果蛇还没死
			if game.getControl().getActivity() {
				//进行渲染界面
				if err := g.screenApp.start(game); err != nil {
					panic(err.Error())
				}
			}

			//刷新帧
			flush(*game.getScreen().getScore())
		}
	}
}
